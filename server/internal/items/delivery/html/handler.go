package html

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/config"
	"github.com/darkalit/rlzone/server/internal/items"
	"github.com/darkalit/rlzone/server/internal/users"
	"github.com/darkalit/rlzone/server/pkg/auth"
	"github.com/darkalit/rlzone/server/pkg/httpErrors"
)

type Handler struct {
	cfg          *config.Config
	useCase      items.UseCase
	usersUseCase users.UseCase
}

func NewHandler(cfg *config.Config, useCase items.UseCase, usersUseCase users.UseCase) *Handler {
	return &Handler{
		cfg:          cfg,
		useCase:      useCase,
		usersUseCase: usersUseCase,
	}
}

func (h *Handler) Get(c *gin.Context) {
	var user *users.User
	query := items.GetItemsQuery{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	payloadAny, exists := c.Get("payload")
	if exists {
		payload := payloadAny.(auth.JWTPayload)
		user, err = h.usersUseCase.GetById(c.Request.Context(), payload.UserID)
		if err != nil {
			c.JSON(httpErrors.ErrorResponse(err))
		}
	}

	itemsResponse, err := h.useCase.Get(c.Request.Context(), &query)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}
	query.Page = int(itemsResponse.Pagination.Page)
	query.PageSize = int(itemsResponse.Pagination.Size)

	c.HTML(http.StatusOK, "items.html", gin.H{
		"items":      itemsResponse.Items,
		"query":      query,
		"pagination": itemsResponse.Pagination,
		"user":       user,
	})
}

func (h *Handler) GetList(c *gin.Context) {
	query := items.GetItemsQuery{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}
	query.PageSize = math.MaxInt

	itemsResponse, err := h.useCase.Get(c.Request.Context(), &query)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.HTML(http.StatusOK, "modal-search-result.html", gin.H{
		"items": itemsResponse.Items,
	})
}

func (h *Handler) BuyItem(c *gin.Context) {
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

	_, err = h.useCase.BuyItem(c.Request.Context(), query.ItemID, payload.UserID)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.Redirect(http.StatusFound, "/items")
}

func (h *Handler) SellItem(c *gin.Context) {
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

	_, err = h.useCase.SellItem(c.Request.Context(), query.ItemID, payload.UserID)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.Redirect(http.StatusFound, "/items/inventory")
}

func (h *Handler) Inventory(c *gin.Context) {
	query := items.GetItemsQuery{}
	err := c.ShouldBindQuery(&query)
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
	user, err := h.usersUseCase.GetById(c.Request.Context(), payload.UserID)

	itemsResponse, err := h.useCase.GetInventory(c.Request.Context(), &query, user.ID)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}
	query.Page = int(itemsResponse.Pagination.Page)
	query.PageSize = int(itemsResponse.Pagination.Size)

	c.HTML(http.StatusOK, "items.html", gin.H{
		"inventoryItems": itemsResponse.InventoryItems,
		"query":          query,
		"pagination":     itemsResponse.Pagination,
		"user":           user,
	})
}

func (h *Handler) CreateStockGet(c *gin.Context) {
	var user *users.User

	itemsResponse, err := h.useCase.Get(c.Request.Context(), &items.GetItemsQuery{
		PageSize: 1,
	})
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	payloadAny, exists := c.Get("payload")
	if exists {
		payload := payloadAny.(auth.JWTPayload)
		user, err = h.usersUseCase.GetById(c.Request.Context(), payload.UserID)
		if err != nil {
			c.JSON(httpErrors.ErrorResponse(err))
		}
	}

	if len(itemsResponse.Items) == 0 {
		c.JSON(httpErrors.ErrorResponse(httpErrors.NewRestErrorMessage(500, "No items")))
	}

	c.HTML(http.StatusOK, "items-create.html", gin.H{
		"item": itemsResponse.Items[0],
		"user": user,
	})
}

func (h *Handler) CreateStockPost(c *gin.Context) {
	var createStockRequest items.CreateStockRequest

	err := c.ShouldBind(&createStockRequest)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	_, err = h.useCase.CreateStock(c, &createStockRequest)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	c.Redirect(http.StatusFound, "/items")
}
