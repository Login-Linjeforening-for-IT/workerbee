package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateLocation(c *gin.Context) {
	var location models.NewLocation

	if err := c.ShouldBindBodyWithJSON(&location); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, location, *h.Services.Validate) {
		return
	}

	locationResponse, err := h.Services.Locations.CreateLocation(location)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, locationResponse)
}

func (h *Handler) GetLocations(c *gin.Context) {
	types := c.DefaultQuery("type", "")
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	sort := c.DefaultQuery("sort", "asc")
	orderBy := c.DefaultQuery("order_by", "id")

	locs, err := h.Services.Locations.GetLocations(search, limit, offset, orderBy, sort, types)
	if internal.HandleError(c, err) {
		return
	}

	if len(locs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"locations":   locs,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"locations":   locs,
			"total_count": locs[0].TotalCount,
		})
	}
}

func (h *Handler) GetLocationNames(c *gin.Context) {
	locationNames, err := h.Services.Locations.GetLocationNames()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, locationNames)
}

func (h *Handler) GetLocation(c *gin.Context) {
	id := c.Param("id")

	loc, err := h.Services.Locations.GetLocation(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, loc)
}

func (h *Handler) GetAllLocationTypes(c *gin.Context) {
	locationTypes, err := h.Services.Locations.GetAllLocationTypes()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, locationTypes)
}

func (h *Handler) UpdateLocation(c *gin.Context) {
	var location models.NewLocation
	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&location); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, location, *h.Services.Validate) {
		return
	}

	location, err := h.Services.Locations.UpdateLocation(id, location)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, location)
}

func (h *Handler) DeleteLocation(c *gin.Context) {
	id := c.Param("id")

	locId, err := h.Services.Locations.DeleteLocation(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": locId})
}
