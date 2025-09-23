package handlers

import (
	"log"
	"net/http"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOranizations(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	sort := c.DefaultQuery("sort", "asc")
	orderBy := c.DefaultQuery("order_by", "id")

	orgs, err := h.Services.Organizations.GetOrgs(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"organizations": orgs,
		"count":         orgs[0].TotalCount,
	})
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
