package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAudience(c *gin.Context) {
	id := c.Param("id")

	audience, err := h.Services.Audiences.GetAudience(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, audience)
}

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
