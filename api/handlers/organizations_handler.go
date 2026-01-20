package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

// CreateOrganization godoc
// @Summary      Create a new organization
// @Description  Creates a new organization with the provided details.
// @Tags         organizations
// @Accept       json
// @Produce      json
// @Param        organization  body      models.Organization  true  "Organization to create"
// @Success      201           {object}  models.Organization
// @Failure      400           {object}  error
// @Failure      500           {object}  error
// @Router       /api/v2/organizations [post]
func (h *Handler) CreateOrganization(c *gin.Context) {
	var org models.Organization

	if err := c.ShouldBindBodyWithJSON(&org); internal.HandleError(c, err) {
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

// UpdateOrganization godoc
// @Summary      Update an existing organization
// @Description  Updates an existing organization with the provided details.
// @Tags         organizations
// @Accept       json
// @Produce      json
// @Param        id            path      string              true  "Organization ID"
// @Param        organization  body      models.Organization  true  "Updated organization details"
// @Success      200           {object}  models.Organization
// @Failure      400           {object}  error
// @Failure      500           {object}  error
// @Router       /api/v2/organizations/{id} [put]
func (h *Handler) UpdateOrganization(c *gin.Context) {
	var org models.Organization
	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&org); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, org, *h.Services.Validate) {
		return
	}

	orgResponse, err := h.Services.Organizations.UpdateOrg(id, org)
	if internal.HandleError(c, err) {
		return
	}

	SetSurrogatePurgeHeader(c,
		"organizations",
		"jobs",
	)

	c.JSON(http.StatusOK, orgResponse)
}

// GetOrganizations godoc
// @Summary      Get organizations
// @Description  Retrieves a list of organizations with optional search, pagination, and sorting.
// @Tags         organizations
// @Produce      json
// @Param        search    query     string  false  "Search term"
// @Param        limit     query     string  false  "Number of records to return"  default(20)
// @Param        offset    query     string  false  "Number of records to skip"     default(0)
// @Param        sort      query     string  false  "Sort order (asc or desc)"      default(asc)
// @Param        order_by  query     string  false  "Field to order by"             default(id)
// @Success      200       {object}  map[string]interface{}
// @Failure      500       {object}  error
// @Router       /api/v2/organizations [get]
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

	SetSurrogatePurgeHeader(c,
		"organizations",
	)

	if len(orgs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"organizations": orgs,
			"total_count":   0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"organizations": orgs,
			"total_count":   orgs[0].TotalCount,
		})
	}
}

// GetOrganizationNames godoc
// @Summary      Get organization names
// @Description  Retrieves a list of all organization names.
// @Tags         organizations
// @Produce      json
// @Success      200  {array}   models.OrganizationNames
// @Failure      500  {object}  error
// @Router       /api/v2/organizations/all [get]
func (h *Handler) GetOrganizationNames(c *gin.Context) {
	orgNames, err := h.Services.Organizations.GetOrgNames()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, orgNames)
}

// GetOrganization godoc
// @Summary      Get organization by ID
// @Description  Retrieves a specific organization by its ID.
// @Tags         organizations
// @Produce      json
// @Param        id   path      string  true  "Organization ID"
// @Success      200  {object}  models.Organization
// @Failure      500  {object}  error
// @Router       /api/v2/organizations/{id} [get]
func (h *Handler) GetOrganization(c *gin.Context) {
	id := c.Param("id")

	org, err := h.Services.Organizations.GetOrg(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, org)
}

// GetOrganization godoc
// @Summary      Delete an organization
// @Description  Deletes an organization by ID.
// @Tags         organizations
// @Produce      json
// @Param        id   path      string  true  "Organization ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  error
// @Router       /api/v2/organizations/{id} [delete]
func (h *Handler) DeleteOrganization(c *gin.Context) {
	id := c.Param("id")

	orgId, err := h.Services.Organizations.DeleteOrg(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": orgId})
}
