package handlers

import (
	"log"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAlbum(c *gin.Context) {
	log.Println("CreateAlbum called")

	form, err := c.MultipartForm()
	if internal.HandleError(c, err) {
		return
	}

	images := form.File["images"]
	if len(images) == 0 {
		internal.HandleError(c, internal.ErrNoImagesProvided)
		return
	}

	var body models.Album
	if err := c.ShouldBindBodyWithJSON(&body); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, body, *h.Services.Validate) {
		return
	}

	err = h.Services.Albums.CreateAlbum(c.Request.Context(), images, body)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "album created successfully",
	})
}
