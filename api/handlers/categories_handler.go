package handlers

import (
	"net/http"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCategories(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "desc")

	categories, err := h.Services.Categories.GetCategories(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories":  categories,
		"total_count": categories[0].TotalCount,
	})
}

func (h *Handler) GetCategory(c *gin.Context) {
	id := c.Param("id")

	category, err := h.Services.Categories.GetCategory(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	category, err := h.Services.Categories.DeleteCategory(id)
	if internal.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, category)
}
