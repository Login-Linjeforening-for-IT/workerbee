package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/db"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
)

// GetEvents godoc
// @Summary      Get events
// @Description  Returns a list of events with details, including category, location, audience, and organizer info. Supports historical filtering, limit, and offset.
// @Tags         events
// @Produce      json
// @Param        limit    query  int   false  "Maximum number of results"
// @Param        offset   query  int   false  "Offset for pagination"
// @Param        historical query bool false  "Include historical events"
// @Success      200  {array}  models.Event
// @Failure      500  {object}  error
// @Router       /api/v2/events [get]
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
