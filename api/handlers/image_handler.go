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
