package handlers

import (
	"log"
	"net/http"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetLocations(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	sort := c.DefaultQuery("sort", "asc")
	orderBy := c.DefaultQuery("order_by", "id")

	orgs, err := h.Services.Locations.GetLocations(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"organizations": orgs,
		"count":         orgs[0].TotalCount,
	})
}
