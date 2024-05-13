package items

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/config"
	"github.com/darkalit/rlzone/server/pkg/httpErrors"
)

type Handler struct {
	cfg     *config.Config
	useCase UseCase
}

func NewHandler(cfg *config.Config, useCase UseCase) *Handler {
	return &Handler{
		cfg:     cfg,
		useCase: useCase,
	}
}

func (h *Handler) Get(c *gin.Context) {
	query := GetItemsQuery{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	itemsResponse, err := h.useCase.Get(c.Request.Context(), &query)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, itemsResponse)
}

func (h *Handler) GetById(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	id := uint(id64)
	fullItem, err := h.useCase.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, fullItem)
}

func (h *Handler) CreateStock(c *gin.Context) {
	var createStockRequest CreateStockRequest
	err := c.ShouldBindJSON(&createStockRequest)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	createdStock, err := h.useCase.CreateStock(c, &createStockRequest)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, createdStock)
}
