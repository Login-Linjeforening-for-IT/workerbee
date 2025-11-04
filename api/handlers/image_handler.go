package handlers

import (
	"net/http"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UploadImage(c *gin.Context) {
	path := c.Param("path")

	file, err := c.FormFile("image")
	if internal.HandleError(c, err) {
		return
	}

	imageURL, err := h.Services.ImageService.UploadImage(file, c.Request.Context(), path)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, imageURL)
}

func (h *Handler) GetImageURLs(c *gin.Context) {
	path := c.Param("path")

	imageURLs, err := h.Services.ImageService.GetImagesInPath(c.Request.Context(), path)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, imageURLs)
}

func (h *Handler) DeleteImage(c *gin.Context) {
	path := c.Param("path")
	imageName := c.Param("imageName")

	path, err := h.Services.ImageService.DeleteImage(c.Request.Context(), path, imageName)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, path)
}
