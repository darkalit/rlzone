package items

import (
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
	query := GetItemQuery{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	items, err := h.useCase.Get(c.Request.Context(), &query)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(200, gin.H{
		"data": items,
	})
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

	c.JSON(200, fullItem)
}
