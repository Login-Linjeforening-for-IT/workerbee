package handlers

import (
	"log"
	"workerbee/internal"

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
}
