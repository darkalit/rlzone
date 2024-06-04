package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/darkalit/rlzone/server/config"
	"github.com/darkalit/rlzone/server/internal/items"
	"github.com/darkalit/rlzone/server/internal/users"
	restUsers "github.com/darkalit/rlzone/server/internal/users/delivery/rest"
	"github.com/darkalit/rlzone/server/pkg/db/mysql"
)

type ItemsHandlerSuite struct {
	suite.Suite
	usecase       *items.ItemsUseCase
	handler       *Handler
	testingServer *httptest.Server
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
	db.Migrator().DropTable(&items.Stock{})
	db.Migrator().CreateTable(&items.Stock{})
	db.Migrator().DropTable(&items.InventoryItem{})
	db.Migrator().CreateTable(&items.InventoryItem{})
	db.Migrator().DropTable(&users.User{})
	db.Migrator().CreateTable(&users.User{})

	repo := items.NewItemRepository(cfg, db)
	usecase := items.NewItemUseCase(repo)
	handler := NewHandler(cfg, usecase)

	usersRepo := users.NewUserRepository(cfg, db)
	usersUseCase := users.NewUserUseCase(usersRepo, cfg)
	usersHandler := restUsers.NewHandler(cfg, usersUseCase)

	router := gin.Default()
	router.GET("/items", handler.Get)
	router.GET("/items/:id", handler.GetById)
	router.POST("/items/stocks", handler.CreateStock)
	router.POST("/users/register", usersHandler.Register)

	s.usecase = usecase
	s.handler = handler
	s.testingServer = httptest.NewServer(router)

	registerRequest := users.RegisterRequest{
		EpicID:   "admin",
		Password: "admin",
		Email:    "admin@gmail.com",
	}
	registerRequestJSON, err := json.Marshal(registerRequest)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post(
		fmt.Sprintf("%s/users/register", s.testingServer.URL),
		"application/json",
		bytes.NewBuffer(registerRequestJSON),
	)
	if err != nil {
		log.Fatal(err)
	}
	resBody := users.UserWithTokens{}
	json.NewDecoder(res.Body).Decode(&resBody)

	db.Model(&users.User{}).Where("id = ?", resBody.User.ID).Update("role", users.RoleAdmin)
}

func (s *ItemsHandlerSuite) TearDownSuite() {
	defer s.testingServer.Close()
}

func (s *ItemsHandlerSuite) TestItemsRestHandler_Get() {
	res, err := http.Get(fmt.Sprintf("%s/items", s.testingServer.URL))
	s.NoError(err, "no error when calling endpoint")
	defer res.Body.Close()

	resBody := items.GetResponse{}
	json.NewDecoder(res.Body).Decode(&resBody)

	s.Equal(http.StatusOK, res.StatusCode)
}

func (s *ItemsHandlerSuite) TestItemsRestHandler_GetById() {
	var itemId uint = 32

	res, err := http.Get(fmt.Sprintf("%s/items/%d", s.testingServer.URL, itemId))
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

	res, err := http.Post(
		fmt.Sprintf("%s/items/stocks", s.testingServer.URL),
		"application/json",
		bytes.NewBuffer(createStockRequestJSON),
	)
	log.Printf("FFFFFFFFFFFFFFFFFFUCK %+v", res.Body)
	s.NoError(err, "no error when calling endpoint")
	defer res.Body.Close()

	resBodyStock := items.Stock{}
	json.NewDecoder(res.Body).Decode(&resBodyStock)

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal(createStockRequest.ItemID, resBodyStock.ItemID)

	res, err = http.Get(fmt.Sprintf("%s/items/%d", s.testingServer.URL, createStockRequest.ItemID))
	s.NoError(err, "no error when calling endpoint")
	defer res.Body.Close()

	resBodyItem := items.Item{}
	json.NewDecoder(res.Body).Decode(&resBodyItem)

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal(resBodyStock.ID, resBodyItem.Stock.ID)
}
