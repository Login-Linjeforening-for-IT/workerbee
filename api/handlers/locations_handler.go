package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

// CreateLocation godoc
// @Summary      Create a new location
// @Description  Creates a new location with the provided details.
// @Tags         locations
// @Accept       json
// @Produce      json
// @Param        location  body      models.NewLocation  true  "Location to create"
// @Success      201       {object}  models.Location
// @Failure      400       {object}  error
// @Failure      500       {object}  error
// @Router       /api/v2/locations [post]
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

// GetLocations godoc
// @Summary      Get locations
// @Description  Retrieves a list of locations with optional filtering, sorting, and pagination.
// @Tags         locations
// @Produce      json
// @Param        type      query     string  false  "Filter by location type"
// @Param        search    query     string  false  "Search term"
// @Param        limit     query     string  false  "Number of records to return"  default(20)
// @Param        offset    query     string  false  "Number of records to skip"    default(0)
// @Param        sort      query     string  false  "Sort order (asc or desc)"    default(asc)
// @Param        order_by  query     string  false  "Field to order by"           default(id)
// @Success      200       {object}  map[string]interface{}
// @Failure      500       {object}  error
// @Router       /api/v2/locations [get]
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

// GetLocationNames godoc
// @Summary      Get location names
// @Description  Retrieves a list of all location names.
// @Tags         locations
// @Produce      json
// @Success      200  {array}   string
// @Failure      500  {object}  error
// @Router       /api/v2/locations/names [get]
func (h *Handler) GetLocationNames(c *gin.Context) {
	locationNames, err := h.Services.Locations.GetLocationNames()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, locationNames)
}

// GetLocation godoc
// @Summary      Get location by ID
// @Description  Retrieves a specific location by its ID.
// @Tags         locations
// @Produce      json
// @Param        id   path      string  true  "Location ID"
// @Success      200  {object}  models.Location
// @Failure      500  {object}  error
// @Router       /api/v2/locations/{id} [get]
func (h *Handler) GetLocation(c *gin.Context) {
	id := c.Param("id")

	loc, err := h.Services.Locations.GetLocation(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, loc)
}

// GetAllLocationTypes godoc
// @Summary      Get all location types
// @Description  Retrieves a list of all location types.
// @Tags         locations
// @Produce      json
// @Success      200  {array}   string
// @Failure      500  {object}  error
// @Router       /api/v2/locations/types [get]
func (h *Handler) GetAllLocationTypes(c *gin.Context) {
	locationTypes, err := h.Services.Locations.GetAllLocationTypes()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, locationTypes)
}

// UpdateLocation godoc
// @Summary      Update a location
// @Description  Updates a location by ID with the provided details.
// @Tags         locations
// @Accept       json
// @Produce      json
// @Param        id        path      string             true  "Location ID"
// @Param        location  body      models.NewLocation  true  "Location object"
// @Success      200       {object}  models.Location
// @Failure      400       {object}  error
// @Failure      500       {object}  error
// @Router       /api/v2/locations/{id} [put]
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

// DeleteLocation godoc
// @Summary      Delete a location
// @Description  Deletes a location by ID.
// @Tags         locations
// @Produce      json
// @Param        id   path      string  true  "Location ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  error
// @Router       /api/v2/locations/{id} [delete]
func (h *Handler) DeleteLocation(c *gin.Context) {
	id := c.Param("id")

	locId, err := h.Services.Locations.DeleteLocation(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": locId})
}
