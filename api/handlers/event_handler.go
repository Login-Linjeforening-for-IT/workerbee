package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateEvent(c *gin.Context) {
	var event models.NewEvent

	if err := c.ShouldBindBodyWithJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body",
		})
		return
	}

	if internal.HandleValidationError(c, event, *h.Services.Validate) {
		return
	}

	event, err := h.Services.Events.CreateEvent(event)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, event)
}

func (h *Handler) UpdateEvent(c *gin.Context) {
	var event models.NewEvent
	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body",
		})
		return
	}

	if internal.HandleValidationError(c, event, *h.Services.Validate) {
		return
	}

	event, err := h.Services.Events.UpdateEvent(event, id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, event)
}

func (h *Handler) GetEvents(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	categories := c.DefaultQuery("categories", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")
	historical := c.DefaultQuery("historical", "false")

	events, err := h.Services.Events.GetEvents(search, limit, offset, orderBy, sort, historical, categories)
	if internal.HandleError(c, err) {
		return
	}

	if len(events) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"events":      events,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"events":      events,
			"total_count": events[0].TotalCount,
		})
	}
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
func (h *Handler) GetAllEvents(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	categories := c.DefaultQuery("categories", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")
	historical := c.DefaultQuery("historical", "false")

	events, err := h.Services.Events.GetAllEvents(search, limit, offset, orderBy, sort, historical, categories)
	if internal.HandleError(c, err) {
		return
	}

	if len(events) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"events":      events,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"events":      events,
			"total_count": events[0].TotalCount,
		})
	}
}

func (h *Handler) GetEvent(c *gin.Context) {
	id := c.Param("id")

	event, err := h.Services.Events.GetEvent(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, event)
}

func (h *Handler) GetAllEventCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"categories": h.Services.Events.GetAllEventCategories(),
	})
}

func (h *Handler) GetEventAudiences(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"audiences": h.Services.Events.GetEventAudiences(),
	})
}

func (h *Handler) GetEventCategories(c *gin.Context) {
	categories, err := h.Services.Events.GetEventCategories()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func (h *Handler) DeleteEvent(c *gin.Context) {
	id := c.Param("id")

	eventId, err := h.Services.Events.DeleteEvent(id)
	if internal.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": eventId})
}
