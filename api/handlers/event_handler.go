package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/db"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
)

func GetEvents(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	println(limit, offset)

	events := []models.Event{}

	sqlBytes, err := os.ReadFile("./db/events/get_events.sql")
	if err != nil {
		log.Println("unable to find file, err ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	historical := true
	query := string(sqlBytes)

	err = db.DB.Select(&events, query, historical, 20, 0)
	if err != nil {
		log.Println("unable to query, err ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, events)
}
