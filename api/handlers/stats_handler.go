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
// @Success      200  {object}  []models.YearlyActivity
// @Failure      500  {object}  error
// @Router       /api/v2/stats/total [get]
func (h *Handler) GetYearlyStats(c *gin.Context) {
	yearlyStats, err := h.Services.Stats.GetYearlyStats()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, yearlyStats)
}

// GetCategoriesStats godoc
// @Summary      Get category event statistics
// @Description  Returns, for each category, the number of events in the last 3 months. Only categories with at least one event are included, ordered by event count descending.
// @Tags         stats
// @Produce      json
// @Success      200  {array}  map[string]interface{}
// @Failure      500  {object}  error
// @Router       /api/v2/stats/categories [get]
func (h *Handler) GetMostActiveCategories(c *gin.Context) {
	categoriesStats, err := h.Services.Stats.GetMostActiveCategories()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, categoriesStats)
}

// GetNewAdditionsStats godoc
// @Summary      Get newest additions
// @Description  Returns a list of the newest additions across events, categories, audiences, rules, organizations, locations, and job advertisements. Ordered by creation date, limited by the 'limit' query parameter (default 10, max 25).
// @Tags         stats
// @Produce      json
// @Param        limit  query  int  false  "Maximum number of results (1-25)"
// @Success      200  {array}  models.NewAddition
// @Failure      500  {object}  error
// @Router       /api/v2/stats/new-additions [get]
func (h *Handler) GetNewAdditionsStats(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")

	NewAdditionsStats, err := h.Services.Stats.GetNewAdditionsStats(limit)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, NewAdditionsStats)
}
