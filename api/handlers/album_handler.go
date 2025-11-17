package handlers

import (
	"log"
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

// CreateAlbum godoc
// @Summary      Create a new album
// @Description  Creates a new album with the provided details
// @Tags         albums
// @Accept       json
// @Produce      json
// @Param        album body      models.CreateAlbum true "Album details"
// @Success      200  {object}  models.Album
// @Failure      400  {object}  error
// @Security     Bearer
// @Router       /api/v2/albums [post]
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

// UploadImagesToAlbum godoc
// @Summary      Upload images to an album
// @Description  Uploads one or more images to the specified album
// @Tags         albums
// @Accept       multipart/form-data
// @Produce      json
// @Param        id     path      string  true  "Album ID"
// @Param        images formData  []file  true  "Images to upload" collectionFormat(multi)
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  error
// @Security     Bearer
// @Router       /api/v2/albums/{id}/images [post]
func (h *Handler) UploadImagesToAlbum(c *gin.Context) {
	id := c.Param("id")

	var files []models.UploadImages
	if err := c.ShouldBindBodyWithJSON(&files); internal.HandleError(c, err) {
		return
	}

	responses, err := h.Services.Albums.UploadImagesToAlbum(c.Request.Context(), id, files)
	if internal.HandleError(c, err) {
		return
	}

	log.Println(responses)

	c.JSON(http.StatusOK, responses)
}

// GetAlbums godoc
// @Summary      Get a list of albums
// @Description  Retrieves a list of albums with optional sorting, pagination, and search
// @Tags         albums
// @Accept       json
// @Produce      json
// @Param        order_by  query     string  false  "Column to order by"  default(id)
// @Param        sort      query     string  false  "Sort direction (asc or desc)"  default(asc)
// @Param        limit     query     string  false  "Number of records to return"  default(20)
// @Param        offset    query     string  false  "Offset for pagination"  default(0)
// @Param        search    query     string  false  "Search term"
// @Success      200  {array}  models.AlbumsWithTotalCount
// @Failure      400  {object}  error
// @Router       /api/v2/albums [get]
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

// GetAlbum godoc
// @Summary      Get album details
// @Description  Retrieves details of a specific album by ID
// @Tags         albums
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Album ID"
// @Success      200  {object}  models.AlbumWithImages
// @Failure      400  {object}  error
// @Router       /api/v2/albums/{id} [get]
func (h *Handler) GetAlbum(c *gin.Context) {
	id := c.Param("id")

	album, err := h.Services.Albums.GetAlbum(c.Request.Context(), id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, album)
}

// UpdateAlbum godoc
// @Summary      Update album details
// @Description  Updates the details of a specific album by ID
// @Tags         albums
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Album ID"
// @Param        album body      models.CreateAlbum true "Updated album details"
// @Success      200  {object}  models.CreateAlbum
// @Failure      400  {object}  error
// @Security     Bearer
// @Router       /api/v2/albums/{id} [put]
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

// DeleteAlbum godoc
// @Summary      Delete an album
// @Description  Deletes a specific album by ID
// @Tags         albums
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Album ID"
// @Success      200  {object}  map[string]int
// @Failure      400  {object}  error
// @Security     Bearer
// @Router       /api/v2/albums/{id} [delete]
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

// DeleteAlbumImage godoc
// @Summary      Delete an image from an album
// @Description  Deletes a specific image from an album by album ID and image name
// @Tags         albums
// @Accept       json
// @Produce      json
// @Param        id        path      string  true  "Album ID"
// @Param        imageName path      string  true  "Image Name"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  error
// @Security     Bearer
// @Router       /api/v2/albums/{id}/images/{imageName} [delete]
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

// SetAlbumCover godoc
// @Summary      Set album cover image
// @Description  Sets a specific image as the cover for an album by album ID and image name
// @Tags         albums
// @Accept       json
// @Produce      json
// @Param        id        path      string  true  "Album ID"
// @Param        imageName path      string  true  "Image Name"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  error
// @Security     Bearer
// @Router       /api/v2/albums/{id}/cover/{imageName} [put]
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
