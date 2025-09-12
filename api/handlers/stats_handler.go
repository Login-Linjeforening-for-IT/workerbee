package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/db"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
)

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

