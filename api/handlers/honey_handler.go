package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

// CreateTextInService godoc
// @Summary      Create text content in a specified service and path
// @Description  Creates text content in the specified service and path for a given language.
// @Tags         honey
// @Accept       json
// @Produce      json
// @Param        service    path      string                           true  "Service name"
// @Param        path       path      string                           true  "Path in the service"
// @Param        language   path      string                           true  "Language code"
// @Param        content    body      map[string]map[string]string     true  "Content to create"
// @Success      200        {object}  map[string]interface{}
// @Failure      400        {object}  error
// @Failure      500        {object}  error
// @Router       /api/v2/honey/{service}/content/{path}/{language} [post]
func (h *Handler) CreateHoney(c *gin.Context) {
	var content models.CreateHoney
	if err := c.ShouldBindJSON(&content); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, content, *h.Services.Validate) {
		return	
	}

	response, err := h.Services.Honey.CreateHoney(content)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetTextServices godoc
// @Summary      Get available text services
// @Description  Retrieves a list of all available text services.
// @Tags         honey
// @Produce      json
// @Success      200  {array}   string
// @Failure      500  {object}  error
// @Router       /api/v2/honey/services [get]
func (h *Handler) GetTextServices(c *gin.Context) {
	services, err := h.Services.Honey.GetTextServices()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, services)
}

// GetAllPathsInService godoc
// @Summary      Get all paths in a specified service
// @Description  Retrieves all paths available in the specified text service.
// @Tags         honey
// @Produce      json
// @Param        service    path      string  true  "Service name"
// @Success      200        {array}   string
// @Failure      500        {object}  error
// @Router       /api/v2/honey/{service}/paths [get]
func (h *Handler) GetAllPathsInService(c *gin.Context) {
	service := c.Param("service")

	paths, err := h.Services.Honey.GetAllPathsInService(service)
	if internal.HandleError(c, err) {
		return
	}

	if len(paths) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"honeys":      paths,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"honeys":      paths,
			"total_count": paths[0].TotalCount,
		})
	}
}

// GetAllContentInPath godoc
// @Summary      Get all content in a specified path
// @Description  Retrieves all text content available in the specified path of a text service.
// @Tags         honey
// @Produce      json
// @Param        service    path      string  true  "Service name"
// @Param        path       path      string  true  "Path in the service"
// @Success      200        {object}  map[string]map[string]string
// @Failure      500        {object}  error
// @Router       /api/v2/honey/{service}/content/{path} [get]
func (h *Handler) GetAllContentInPath(c *gin.Context) {
	service := c.Param("service")
	path := c.Param("path")

	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	content, err := h.Services.Honey.GetAllContentInPath(service, path)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, content)
}

// GetOneLanguage godoc
// @Summary      Get content for one language in a specified path
// @Description  Retrieves text content for a specific language in the specified path of a text service.
// @Tags         honey
// @Produce      json
// @Param        service    path      string  true  "Service name"
// @Param        path       path      string  true  "Path in the service"
// @Param        language   path      string  true  "Language code"
// @Success      200        {object}  map[string]string
// @Failure      500        {object}  error
// @Router       /api/v2/honey/{service}/content/{path}/{language} [get]
func (h *Handler) GetOneLanguage(c *gin.Context) {
	service := c.Param("service")
	path := c.Param("path")
	language := c.Param("language")

	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	response, err := h.Services.Honey.GetOneLanguage(service, path, language)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetHoney godoc
// @Summary      Get honey by ID
// @Description  Retrieves a honey entry by its ID
// @Tags         honey
// @Produce      json
// @Param        id   path      int  true  "Honey ID"
// @Success      200  {object}  models.CreateHoney
// @Failure      400  {object}  error
// @Router       /api/v2/honey/{id} [get]
func (h *Handler) GetHoney(c *gin.Context) {
	id := c.Param("id")

	honey, err := h.Services.Honey.GetHoney(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, honey)
}

// UpdateContentInPath godoc
// @Summary      Update content in a specified path
// @Description  Updates text content in the specified path of a text service.
// @Tags         honey
// @Accept       json
// @Produce      json
// @Param        service    path      string                           true  "Service name"
// @Param        path       path      string                           true  "Path in the service"
// @Param        content    body      map[string]map[string]string     true  "Content to update"
// @Success      200        {object}  map[string]interface{}
// @Failure      400        {object}  error
// @Failure      500        {object}  error
// @Router       /api/v2/honey/{service}/content/{path} [put]
func (h *Handler) UpdateHoney(c *gin.Context) {
	id := c.Param("id")

	var content models.CreateHoney
	if err := c.ShouldBindJSON(&content); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, content, *h.Services.Validate) {
		return	
	}

	resp, err := h.Services.Honey.UpdateContentInPath(id, content)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteHoney(c *gin.Context) {
	id := c.Param("id")

	honeyID, err := h.Services.Honey.DeleteHoney(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": honeyID,
	})
}
