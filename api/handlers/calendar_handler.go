package handlers

import (
	"net/http"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCalendar(c *gin.Context) {
	categories := c.DefaultQuery("categories", "")
	language := c.DefaultQuery("language", "no")

	cal, err := h.Services.Calendar.GetCalendarData(categories, language)
	if internal.HandleError(c, err) {
		return
	}

	c.Header("Content-Type", "text/calendar; charset=utf-8")
	c.Header("Content-Disposition", "attachment; filename=calendar.ics")
	c.String(http.StatusOK, cal.Serialize())
}
