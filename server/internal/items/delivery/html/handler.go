package html

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/config"
	"github.com/darkalit/rlzone/server/internal/items"
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
	query.Page = int(itemsResponse.Pagination.Page)
	query.PageSize = int(itemsResponse.Pagination.Size)

	c.HTML(http.StatusOK, "items.html", gin.H{
		"items":      itemsResponse.Items,
		"query":      query,
		"pagination": itemsResponse.Pagination,
	})
}
