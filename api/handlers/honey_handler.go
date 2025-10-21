package handlers

import (
	"net/http"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTextServices(c *gin.Context) {
	services, err := h.Services.Honey.GetTextServices()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, services)
}

func (h *Handler) GetAllPathsInService(c *gin.Context) {
	service := c.Param("service")

	paths, err := h.Services.Honey.GetAllPathsInService(service)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, paths)
}

func (h *Handler) GetAllContentInPath(c *gin.Context) {
	service := c.Param("service")
	path := c.Param("path")

	content, err := h.Services.Honey.GetAllContentInPath(service, path)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, content)
}

func (h *Handler) GetOneLanguage(c *gin.Context) {
	service := c.Param("service")
	path := c.Param("path")
	language := c.Param("language")

	response, err := h.Services.Honey.GetOneLanguage(service, path, language)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, response)
}
