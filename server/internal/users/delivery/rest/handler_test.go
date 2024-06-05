package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/darkalit/rlzone/server/internal/users"
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
}

func (s *UsersHandlerSuite) TearDownSuite() {
	defer s.testingServer.Close()
}
