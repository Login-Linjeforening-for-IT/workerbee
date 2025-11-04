package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

// GetCategory godoc
// @Summary      Get a single category
// @Description  Returns a category by its ID
// @Tags         categories
// @Param        id   path      string  true  "Category ID"
// @Success      200  {object}  models.Category
// @Failure      500  {object}  error
// @Router       /api/v2/categories/{id} [get]
func (h *Handler) GetCategory(c *gin.Context) {
	id := c.Param("id")

	category, err := h.Services.Categories.GetCategory(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, category)
}

// GetCategories godoc
// @Summary      List categories
// @Description  Returns a list of categories, supports search and pagination
// @Tags         categories
// @Param        search    query     string  false  "Search string"
// @Param        limit     query     int     false  "Limit"
// @Param        offset    query     int     false  "Offset"
// @Param        order_by  query     string  false  "Order by field"
// @Param        sort      query     string  false  "Sort order (asc/desc)"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  error
// @Router       /api/v2/categories [get]
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

// CreateCategory godoc
// @Summary      Create a new category
// @Description  Creates a new category with the provided details
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category  body      models.Category  true  "Category object"
// @Success      201  {object}  models.Category
// @Failure      400  {object}  error
// @Router       /api/v2/categories [post]
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

// UpdateCategory godoc
// @Summary      Update category details
// @Description  Updates the details of a specific category by ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Category ID"
// @Param        category body      models.Category true "Updated category details"
// @Success      200  {object}  models.Category
// @Failure      400  {object}  error
// @Router       /api/v2/categories/{id} [put]
func (h *Handler) UpdateCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&category); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, category, *h.Services.Validate) {
		return
	}

	categoryResponse, err := h.Services.Categories.UpdateCategory(id, category)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, categoryResponse)
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	catID, err := h.Services.Categories.DeleteCategory(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": catID})
}
