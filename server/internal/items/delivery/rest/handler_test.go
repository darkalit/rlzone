package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/darkalit/rlzone/server/config"
	"github.com/darkalit/rlzone/server/internal/items"
	"github.com/darkalit/rlzone/server/internal/middleware"
	"github.com/darkalit/rlzone/server/internal/users"
	restUsers "github.com/darkalit/rlzone/server/internal/users/delivery/rest"
	"github.com/darkalit/rlzone/server/pkg/db/mysql"
)

type ItemsHandlerSuite struct {
	suite.Suite
	usecase       *items.ItemsUseCase
	handler       *Handler
	testingServer *httptest.Server
	testingClient *http.Client
}

func TestItemsHandlerSuite(t *testing.T) {
	suite.Run(t, new(ItemsHandlerSuite))
}

func (s *ItemsHandlerSuite) SetupSuite() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	cfg.DBName = "rlzonetest"

	db, err := mysql.NewMySqlDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	tx := db.WithContext(ctx)

	tx.Migrator().DropTable(&items.Stock{})
	tx.Migrator().CreateTable(&items.Stock{})
	tx.Migrator().DropTable(&items.InventoryItem{})
	tx.Migrator().CreateTable(&items.InventoryItem{})
	tx.Migrator().DropTable(&users.User{})
	tx.Migrator().CreateTable(&users.User{})

	repo := items.NewItemRepository(cfg, db)
	usecase := items.NewItemUseCase(repo)
	handler := NewHandler(cfg, usecase)

	usersRepo := users.NewUserRepository(cfg, db)
	usersUseCase := users.NewUserUseCase(usersRepo, cfg)
	usersHandler := restUsers.NewHandler(cfg, usersUseCase)

	mw := middleware.NewMiddlewareManager(cfg, usersUseCase)

	router := gin.Default()
	router.Use(mw.SetPayload)
	router.GET("/items", handler.Get)
	router.GET("/items/:id", handler.GetById)
	router.POST("/items/stocks", handler.CreateStock)
	router.POST("/items/buy", handler.Buy)
	router.POST("/items/sell", handler.Sell)
	router.POST("/users/register", usersHandler.Register)

	s.usecase = usecase
	s.handler = handler
	s.testingServer = httptest.NewServer(router)

	jar, _ := cookiejar.New(nil)
	s.testingClient = &http.Client{
		Jar: jar,
	}

	registerRequest := users.RegisterRequest{
		EpicID:   "admin",
		Password: "admin",
		Email:    "admin@gmail.com",
	}
	registerRequestJSON, err := json.Marshal(registerRequest)
	if err != nil {
		log.Fatal(err)
	}

	res, err := s.testingClient.Post(
		fmt.Sprintf("%s/users/register", s.testingServer.URL),
		"application/json",
		bytes.NewBuffer(registerRequestJSON),
	)
	if err != nil {
		log.Fatal(err)
	}
	resBody := users.UserWithTokens{}
	json.NewDecoder(res.Body).Decode(&resBody)

	tx.Model(&users.User{}).
		Where("id = ?", resBody.User.ID).
		Updates(map[string]interface{}{"role": users.RoleAdmin, "balance": 9999})
}

func (s *ItemsHandlerSuite) TearDownSuite() {
	defer s.testingServer.Close()
}

func (s *ItemsHandlerSuite) TestItemsRestHandler_Get() {
	res, err := s.testingClient.Get(fmt.Sprintf("%s/items", s.testingServer.URL))
	s.NoError(err, "no error when calling endpoint")
	defer res.Body.Close()

	resBody := items.GetResponse{}
	json.NewDecoder(res.Body).Decode(&resBody)

	s.Equal(http.StatusOK, res.StatusCode)
}

func (s *ItemsHandlerSuite) TestItemsRestHandler_GetById() {
	var itemId uint = 52

	res, err := s.testingClient.Get(fmt.Sprintf("%s/items/%d", s.testingServer.URL, itemId))
	s.NoError(err, "no error when calling endpoint")
	defer res.Body.Close()

	resBody := items.Item{}
	json.NewDecoder(res.Body).Decode(&resBody)

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal(itemId, resBody.ID)
}

func (s *ItemsHandlerSuite) TestItemsRestHandler_CreateStock() {
	createStockRequest := items.CreateStockRequest{
		Description: "my description",
		Price:       900,
		ItemID:      32,
	}

	createStockRequestJSON, err := json.Marshal(createStockRequest)
	s.NoError(err, "no error when marshal request")

	res, err := s.testingClient.Post(
		fmt.Sprintf("%s/items/stocks", s.testingServer.URL),
		"application/json",
		bytes.NewBuffer(createStockRequestJSON),
	)
	s.NoError(err, "no error when calling endpoint")
	defer res.Body.Close()

	resBodyStock := items.Stock{}
	json.NewDecoder(res.Body).Decode(&resBodyStock)

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal(createStockRequest.ItemID, resBodyStock.ItemID)

	res, err = s.testingClient.Get(
		fmt.Sprintf("%s/items/%d", s.testingServer.URL, createStockRequest.ItemID),
	)
	s.NoError(err, "no error when calling endpoint")
	defer res.Body.Close()

	resBodyItem := items.Item{}
	json.NewDecoder(res.Body).Decode(&resBodyItem)

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal(resBodyStock.ID, resBodyItem.Stock.ID)
}

func (s *ItemsHandlerSuite) TestItemsRestHandler_Buy() {
	var itemId uint = 69
	createStockRequest := items.CreateStockRequest{
		Description: "my description",
		Price:       900,
		ItemID:      itemId,
	}

	createStockRequestJSON, err := json.Marshal(createStockRequest)
	s.NoError(err, "no error when marshal request")

	res, err := s.testingClient.Post(
		fmt.Sprintf("%s/items/stocks", s.testingServer.URL),
		"application/json",
		bytes.NewBuffer(createStockRequestJSON),
	)
	s.NoError(err, "no error when calling endpoint")
	defer res.Body.Close()

	resBodyStock := items.Stock{}
	json.NewDecoder(res.Body).Decode(&resBodyStock)

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal(createStockRequest.ItemID, resBodyStock.ItemID)

	buyItemRequest := items.BuySellItemRequest{
		ItemID: itemId,
	}

	buyItemRequestJSON, err := json.Marshal(buyItemRequest)
	s.NoError(err, "no error when marshal request")

	res, err = s.testingClient.Post(
		fmt.Sprintf("%s/items/buy", s.testingServer.URL),
		"application/json",
		bytes.NewBuffer(buyItemRequestJSON),
	)
	s.NoError(err, "no error when calling endpoint")
	defer res.Body.Close()

	resBodyInvItem := items.InventoryItem{}
	json.NewDecoder(res.Body).Decode(&resBodyInvItem)

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal(buyItemRequest.ItemID, resBodyInvItem.ItemID)
}

func (s *ItemsHandlerSuite) TestItemsRestHandler_Sell() {
	var itemId uint = 71
	createStockRequest := items.CreateStockRequest{
		Description: "my description",
		Price:       900,
		ItemID:      itemId,
	}

	createStockRequestJSON, err := json.Marshal(createStockRequest)
	s.NoError(err, "no error when marshal request")

	res, err := s.testingClient.Post(
		fmt.Sprintf("%s/items/stocks", s.testingServer.URL),
		"application/json",
		bytes.NewBuffer(createStockRequestJSON),
	)
	s.NoError(err, "no error when calling endpoint")
	defer res.Body.Close()

	resBodyStock := items.Stock{}
	json.NewDecoder(res.Body).Decode(&resBodyStock)

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal(createStockRequest.ItemID, resBodyStock.ItemID)

	buyItemRequest := items.BuySellItemRequest{
		ItemID: itemId,
	}

	buyItemRequestJSON, err := json.Marshal(buyItemRequest)
	s.NoError(err, "no error when marshal request")

	res, err = s.testingClient.Post(
		fmt.Sprintf("%s/items/buy", s.testingServer.URL),
		"application/json",
		bytes.NewBuffer(buyItemRequestJSON),
	)
	s.NoError(err, "no error when calling endpoint")
	defer res.Body.Close()

	resBodyInvItem := items.InventoryItem{}
	json.NewDecoder(res.Body).Decode(&resBodyInvItem)

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal(buyItemRequest.ItemID, resBodyInvItem.ItemID)

	sellItemRequest := items.BuySellItemRequest{
		ItemID: itemId,
	}

	sellItemRequestJSON, err := json.Marshal(sellItemRequest)
	s.NoError(err, "no error when marshal request")

	res, err = s.testingClient.Post(
		fmt.Sprintf("%s/items/sell", s.testingServer.URL),
		"application/json",
		bytes.NewBuffer(sellItemRequestJSON),
	)
	s.NoError(err, "no error when calling endpoint")
	defer res.Body.Close()

	resBodyInvItem = items.InventoryItem{}
	json.NewDecoder(res.Body).Decode(&resBodyInvItem)

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal(sellItemRequest.ItemID, resBodyInvItem.ItemID)
}
