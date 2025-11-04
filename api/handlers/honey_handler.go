package handlers

import (
	"net/http"
	"workerbee/internal"

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
func (h *Handler) CreateTextInService(c *gin.Context) {
	service := c.Param("service")
	path := c.Param("path")
	language := c.Param("language")

	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	var content map[string]map[string]string
	if err := c.ShouldBindJSON(&content); internal.HandleError(c, err) {
		return
	}

	response, err := h.Services.Honey.CreateTextInService(service, path, language, content)
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

	c.JSON(http.StatusOK, paths)
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
func (h *Handler) UpdateContentInPath(c *gin.Context) {
	service := c.Param("service")
	path := c.Param("path")

	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	var content map[string]map[string]string
	if err := c.ShouldBindJSON(&content); internal.HandleError(c, err) {
		return
	}

	resp, err := h.Services.Honey.UpdateContentInPath(service, path, content)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, resp)
}
