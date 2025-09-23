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

	locs, err := h.Services.Locations.GetLocations(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"locations": locs,
		"count":     locs[0].TotalCount,
	})
}

func (h *Handler) GetLocation(c *gin.Context) {
	id := c.Param("id")

	loc, err := h.Services.Locations.GetLocation(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, loc)
}

func (h *Handler) DeleteLocation(c *gin.Context) {
	id := c.Param("id")

	loc, err := h.Services.Locations.DeleteLocation(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, loc)
}
