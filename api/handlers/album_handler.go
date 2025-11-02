package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAlbum(c *gin.Context) {
	var body models.CreateAlbum
	if err := c.ShouldBindBodyWithJSON(&body); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, body, *h.Services.Validate) {
		return
	}

	albumResponse, err := h.Services.Albums.CreateAlbum(c.Request.Context(), body)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, albumResponse)
}

func (h *Handler) UploadImagesToAlbum(c *gin.Context) {
	id := c.Param("id")

	form, err := c.MultipartForm()
	if internal.HandleError(c, err) {
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		internal.HandleError(c, internal.ErrNoImagesProvided)
		return
	}

	err = h.Services.Albums.UploadImagesToAlbum(c.Request.Context(), id, files)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": "Images have been uploaded",
	})
}

func (h *Handler) GetAlbums(c *gin.Context) {
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	search := c.DefaultQuery("search", "")

	albums, err := h.Services.Albums.GetAlbums(c.Request.Context(), orderBy, sort, limit, offset, search)
	if internal.HandleError(c, err) {
		return
	}

	if len(albums) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"albums":      albums,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"albums":      albums,
			"total_count": albums[0].TotalCount,
		})
	}
}

func (h *Handler) GetAlbum(c *gin.Context) {
	id := c.Param("id")

	album, err := h.Services.Albums.GetAlbum(c.Request.Context(), id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, album)
}

func (h *Handler) UpdateAlbum(c *gin.Context) {
	id := c.Param("id")

	var body models.CreateAlbum
	if err := c.ShouldBindBodyWithJSON(&body); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, body, *h.Services.Validate) {
		return
	}

	albumResponse, err := h.Services.Albums.UpdateAlbum(c.Request.Context(), id, body)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, albumResponse)
}

func (h *Handler) DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	returningID, err := h.Services.Albums.DeleteAlbum(c.Request.Context(), id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": returningID,
	})
}

func (h *Handler) DeleteAlbumImage(c *gin.Context) {
	id := c.Param("id")
	imageName := c.Param("imageName")

	err := h.Services.Albums.DeleteAlbumImage(c.Request.Context(), id, imageName)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Image deleted successfully",
	})
}

func (h *Handler) SetAlbumCover(c *gin.Context) {
	id := c.Param("id")
	imageName := c.Param("imageName")

	err := h.Services.Albums.SetAlbumCover(c.Request.Context(), id, imageName)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Album cover set successfully",
	})
}
