package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/config"
	"github.com/darkalit/rlzone/server/internal/items"
	"github.com/darkalit/rlzone/server/pkg/auth"
	"github.com/darkalit/rlzone/server/pkg/httpErrors"
)

type Handler struct {
	cfg     *config.Config
	useCase items.UseCase
}

func NewHandler(cfg *config.Config, useCase items.UseCase) *Handler {
	return &Handler{
		cfg:     cfg,
		useCase: useCase,
	}
}

func (h *Handler) Get(c *gin.Context) {
	query := items.GetItemsQuery{}
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
	var createStockRequest items.CreateStockRequest
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

func (h *Handler) Buy(c *gin.Context) {
	query := items.BuySellItemRequest{}
	err := c.ShouldBind(&query)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	payloadAny, exists := c.Get("payload")
	if !exists {
		c.JSON(
			httpErrors.ErrorResponse(
				httpErrors.NewRestErrorMessage(http.StatusUnauthorized, "JWT is undefined"),
			),
		)
	}
	payload := payloadAny.(auth.JWTPayload)

	inventoryItem, err := h.useCase.BuyItem(c.Request.Context(), query.ItemID, payload.UserID)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, inventoryItem)
}

func (h *Handler) Sell(c *gin.Context) {
	query := items.BuySellItemRequest{}
	err := c.ShouldBind(&query)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	payloadAny, exists := c.Get("payload")
	if !exists {
		c.JSON(
			httpErrors.ErrorResponse(
				httpErrors.NewRestErrorMessage(http.StatusUnauthorized, "JWT is undefined"),
			),
		)
	}
	payload := payloadAny.(auth.JWTPayload)

	inventoryItem, err := h.useCase.SellItem(c.Request.Context(), query.ItemID, payload.UserID)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, inventoryItem)
}
