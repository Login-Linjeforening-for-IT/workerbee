package handlers

import (
	"net/http"
	"strings"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

var allowedSortColumnsEvents = map[string]string{
	"id":           "e.id",
	"name_no":      "e.name_no",
	"name_en":      "e.name_en",
	"time_start":   "e.time_start",
	"time_end":     "e.time_end",
	"time_publish": "e.time_publish",
	"canceled":     "e.canceled",
	"capacity":     "e.capacity",
	"full":         "e.full",
}

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
func (h *Handler) GetEvents(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	sort := c.DefaultQuery("direction", "asc")
	orderBy := c.DefaultQuery("order_by", "id")
	historical := c.DefaultQuery("historical", "false")

	orderBySanitized, sortSanitized, err := internal.SanitizeSort(orderBy, sort, allowedSortColumnsEvents)
	if internal.HandleError(c, err) {
		return
	}

	events, err := h.Events.GetEvents(search, limit, offset, strings.ToUpper(sortSanitized), orderBySanitized, historical)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"events":      events,
		"total_count": events[0].TotalCount,
	})
}

func (h *Handler) GetEvent(c *gin.Context) {
	id := c.Param("id")

	event, err := h.Events.GetEvent(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, event)
}

func (h *Handler) DeleteEvent(c *gin.Context) {
	id := c.Param("id")

	event, err := h.Events.DeleteEvent(id)
	if internal.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, event)
}
