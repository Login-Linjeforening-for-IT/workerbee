package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

// GetAudience godoc
// @Summary      Get a single audience
// @Description  Returns an audience by its ID
// @Tags         audiences
// @Param        id   path      string  true  "Audience ID"
// @Success      200  {object}  models.Audience
// @Failure      500  {object}  error
// @Router       /api/v2/audiences/{id} [get]
func (h *Handler) GetAudience(c *gin.Context) {
	id := c.Param("id")

	audience, err := h.Services.Audiences.GetAudience(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, audience)
}

// GetAudiences godoc
// @Summary      List audiences
// @Description  Returns a list of audiences, supports search and pagination
// @Tags         audiences
// @Param        search    query     string  false  "Search string"
// @Param        limit     query     int     false  "Limit"
// @Param        offset    query     int     false  "Offset"
// @Param        order_by  query     string  false  "Order by field"
// @Param        sort      query     string  false  "Sort order (asc/desc)"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  error
// @Router       /api/v2/audiences [get]
func (h *Handler) GetAudiences(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")

	audiences, err := h.Services.Audiences.GetAudiences(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		return
	}

	if len(audiences) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"audiences":   audiences,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"audiences":   audiences,
			"total_count": audiences[0].TotalCount,
		})
	}
}

// CreateAudience godoc
// @Summary      Create a new audience
// @Description  Creates a new audience with the provided details
// @Tags         audiences
// @Accept       json
// @Produce      json
// @Param        audience  body      models.Audience  true  "Audience object"
// @Success      201  {object}  models.Audience
// @Failure      400  {object}  error
// @Router       /api/v2/audiences [post]
func (h *Handler) CreateAudience(c *gin.Context) {
	var audience models.Audience

	if err := c.ShouldBindBodyWithJSON(&audience); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, audience, *h.Services.Validate) {
		return
	}

	audienceResponse, err := h.Services.Audiences.CreateAudience(audience)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, audienceResponse)
}

// UpdateAudience godoc
// @Summary      Update an audience
// @Description  Updates an existing audience by its ID
// @Tags         audiences
// @Accept       json
// @Produce      json
// @Param        id        path      string           true  "Audience ID"
// @Param        audience  body      models.Audience  true  "Updated audience object"
// @Success      200  {object}  models.Audience
// @Failure      400  {object}  error
// @Router       /api/v2/audiences/{id} [put]
func (h *Handler) UpdateAudience(c *gin.Context) {
	var audience models.Audience
	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&audience); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, audience, *h.Services.Validate) {
		return
	}

	audienceResponse, err := h.Services.Audiences.UpdateAudience(id, audience)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, audienceResponse)
}

// DeleteAudience godoc
// @Summary      Delete an audience
// @Description  Deletes an audience by its ID
// @Tags         audiences
// @Param        id   path      string  true  "Audience ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  error
// @Router       /api/v2/audiences/{id} [delete]
func (h *Handler) DeleteAudience(c *gin.Context) {
	id := c.Param("id")

	deletedID, err := h.Services.Audiences.DeleteAudience(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": deletedID,
	})
}
