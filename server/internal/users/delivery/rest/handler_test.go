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
	"github.com/darkalit/rlzone/server/internal/middleware"
	"github.com/darkalit/rlzone/server/internal/users"
	"github.com/darkalit/rlzone/server/pkg/db/mysql"
)

type UsersHandlerSuite struct {
	suite.Suite
	usecase       *users.UsersUseCase
	handler       *Handler
	testingServer *httptest.Server
	testingClient *http.Client
}

func TestItemsHandlerSuite(t *testing.T) {
	suite.Run(t, new(UsersHandlerSuite))
}

func (s *UsersHandlerSuite) SetupSuite() {
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

	tx.Migrator().DropTable(&users.User{})
	tx.Migrator().CreateTable(&users.User{})

	repo := users.NewUserRepository(cfg, db)
	usecase := users.NewUserUseCase(repo, cfg)
	handler := NewHandler(cfg, usecase)

	mw := middleware.NewMiddlewareManager(cfg, usecase)

	router := gin.Default()
	router.Use(mw.SetPayload)
	router.POST("/users/register", handler.Register)
	router.GET("/users/block/:id", handler.BlockUser)

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

func (s *UsersHandlerSuite) TearDownSuite() {
	defer s.testingServer.Close()
}

func (s *UsersHandlerSuite) TestUsersHandlerSuite_BlockUser() {
	ctx := context.Background()

	registerRequest := users.RegisterRequest{
		Email:    "user@gmail.com",
		EpicID:   "user",
		Password: "user",
	}
	userWithTokens, err := s.usecase.Register(ctx, &registerRequest)
	s.NoError(err, "no error on creating user")
	s.Equal(false, userWithTokens.User.IsBlocked)

	res, err := s.testingClient.Get(
		fmt.Sprintf("%s/users/block/%d", s.testingServer.URL, userWithTokens.User.ID),
	)
	s.NoError(err, "no error when calling endpoint")
	defer res.Body.Close()

	user, err := s.usecase.GetById(ctx, userWithTokens.User.ID)
	s.Equal(true, user.IsBlocked)
}
