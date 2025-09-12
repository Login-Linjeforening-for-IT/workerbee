package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/db"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
)

// GetTotalStats godoc
// @Summary      Get total statistics
// @Description  Returns total counts for events, jobs, organizations, locations, and rules (excluding deleted records).
// @Tags         stats
// @Produce      json
// @Success      200  {object}  models.TotalStats
// @Failure      500  {object}  error
// @Router       /api/v2/stats/total [get]
func GetTotalStats(c *gin.Context) {
	totalStats := []models.TotalStats{}

	sqlBytes, err := os.ReadFile("./db/stats/get_total_stats.sql")
	if err != nil {
		log.Println("unable to find file, err ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	query := string(sqlBytes)

	err = db.DB.Select(&totalStats, query)
	if err != nil {
		log.Println("unable to query, err ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
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
func GetCategoriesStats(c *gin.Context) {
	categoriesStats := []models.CategoriesStats{}

	sqlBytes, err := os.ReadFile("./db/stats/get_categories_stats.sql")
	if err != nil {
		log.Println("unable to find file, err ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	query := string(sqlBytes)

	err = db.DB.Select(&categoriesStats, query)
	if err != nil {
		log.Println("unable to query, err ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
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
func GetNewAdditionsStats(c *gin.Context) {
	limit := c.Query("limit")

	if limit == "" {
		limit = "10"
	}

	NewAdditionsStats := []models.NewAdditionsStats{}

	sqlBytes, err := os.ReadFile("./db/stats/get_new_additions_stats.sql")
	if err != nil {
		log.Println("unable to find file, err ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	query := string(sqlBytes)

	err = db.DB.Select(&NewAdditionsStats, query, limit)
	if err != nil {
		log.Println("unable to query, err ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, NewAdditionsStats)
}

