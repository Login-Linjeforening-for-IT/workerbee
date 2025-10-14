package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateOrganization(c *gin.Context) {
	var org models.Organization

	if err := c.ShouldBindBodyWithJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body",
		})
		return
	}

	if internal.HandleValidationError(c, org, *h.Services.Validate) {
		return
	}

	orgResponse, err := h.Services.Organizations.CreateOrg(org)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, orgResponse)
}

func (h *Handler) UpdateOrganization(c *gin.Context) {
	var org models.Organization
	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body",
		})
		return
	}

	if internal.HandleValidationError(c, org, *h.Services.Validate) {
		return
	}

	orgResponse, err := h.Services.Organizations.UpdateOrg(id, org)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, orgResponse)
}

func (h *Handler) GetOrganizations(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	sort := c.DefaultQuery("sort", "asc")
	orderBy := c.DefaultQuery("order_by", "id")

	orgs, err := h.Services.Organizations.GetOrgs(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		return
	}

	if len(orgs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"organizations": orgs,
			"total_count":         0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"organizations": orgs,
			"total_count":         orgs[0].TotalCount,
		})
	}
}

func (h *Handler) GetOrganization(c *gin.Context) {
	id := c.Param("id")

	org, err := h.Services.Organizations.GetOrg(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, org)
}

func (h *Handler) DeleteOrganization(c *gin.Context) {
	id := c.Param("id")

	org, err := h.Services.Organizations.DeleteOrg(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, org)
}
