package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCategory(c *gin.Context) {
	id := c.Param("id")

	category, err := h.Services.Categories.GetCategory(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *Handler) GetCategories(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")

	categories, err := h.Services.Categories.GetCategories(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		return
	}

	if len(categories) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"categories":  categories,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"categories":  categories,
			"total_count": categories[0].TotalCount,
		})
	}
}

func (h *Handler) CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindBodyWithJSON(&category); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, category, *h.Services.Validate) {
		return
	}

	categoryResponse, err := h.Services.Categories.CreateCategory(category)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, categoryResponse)
}