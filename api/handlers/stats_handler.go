package handlers

import (
	"net/http"
	"strconv"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

// GetTotalStats godoc
// @Summary      Get total statistics
// @Description  Returns total counts for events, jobs, organizations, locations, and rules (excluding deleted records).
// @Tags         stats
// @Produce      json
// @Success      200  {object}  models.TotalStats
// @Failure      500  {object}  error
// @Router       /api/v2/stats/total [get]
func (h *Handler) GetTotalStats(c *gin.Context) {
	totalStats, err := h.Services.Stats.GetTotalStats()
	if err != nil {

	}

	c.JSON(http.StatusOK, totalStats)
}

// GetCategoriesStats godoc
// @Summary      Get category event statistics
// @Description  Returns, for each category, the number of events in the last 3 months. Only categories with at least one event are included, ordered by event count descending.
// @Tags         stats
// @Produce      json
// @Success      200  {array}  models.CategoriesStats
// @Failure      500  {object}  error
// @Router       /api/v2/stats/categories [get]
func (h *Handler) GetCategoriesStats(c *gin.Context) {
	categoriesStats, err := h.Services.Stats.GetCategoriesStats()
	if err != nil {

	}

	c.JSON(http.StatusOK, categoriesStats)
}

// GetNewAdditionsStats godoc
// @Summary      Get newest additions
// @Description  Returns a list of the newest additions across events, categories, audiences, rules, organizations, locations, and job advertisements. Ordered by creation date, limited by the 'limit' query parameter.
// @Tags         stats
// @Produce      json
// @Param        limit  query  int  false  "Maximum number of results"
// @Success      200  {array}  models.NewAdditionsStats
// @Failure      500  {object}  error
// @Router       /api/v2/stats/new_additions [get]
func (h *Handler) GetNewAdditionsStats(c *gin.Context) {
	limit := c.DefaultQuery("limit", "20")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		internal.HandleError(c,internal.ErrInvalid)
		return
	}

	NewAdditionsStats, err := h.Services.Stats.GetNewAdditionsStats(limitInt)

	c.JSON(http.StatusOK, NewAdditionsStats)
}
