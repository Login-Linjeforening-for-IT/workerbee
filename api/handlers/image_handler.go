package handlers

import (
	"fmt"
	"log"
	"net/http"
	"workerbee/internal"

	"github.com/gin-gonic/gin"

	_ "image/jpeg"
)

// UploadImage godoc
// @Summary      Upload an image to a specified path in a digital ocean bucket
// @Description  Uploads an image file to the specified path in the image service.
// @Tags         images
// @Accept       multipart/form-data
// @Produce      json
// @Param        path   path      string  true  "Path to upload the image"
// @Param        image  formData  file    true  "Image file to upload"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  error
// @Failure      500    {object}  error
// @Router       /api/v2/images/{path} [post]
func (h *Handler) UploadImage(c *gin.Context) {
	path := c.Param("path")

	fmt.Println("Content-Type:", c.Request.Header.Get("Content-Type"))

	file, err := c.FormFile("image")
	log.Println(err)
	if err != nil {
		internal.HandleError(c, err)
		return
	}

	imageURL, err := h.Services.ImageService.UploadImage(file, c.Request.Context(), path)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, imageURL)
}

// GetImageURLs godoc
// @Summary      Get image URLs in a specified path in a digital ocean bucket
// @Description  Retrieves a list of image URLs available in the specified path.
// @Tags         images
// @Produce      json
// @Param        path   path      string  true  "Path to retrieve images from"
// @Success      200    {array}   string
// @Failure      500    {object}  error
// @Router       /api/v2/images/{path} [get]
func (h *Handler) GetImageURLs(c *gin.Context) {
	path := c.Param("path")

	imageURLs, err := h.Services.ImageService.GetImagesInPath(c.Request.Context(), path)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, imageURLs)
}

// DeleteImage godoc
// @Summary      Delete an image from a specified path in a digital ocean bucket
// @Description  Deletes an image file from the specified path in the image service.
// @Tags         images
// @Produce      json
// @Param        path       path      string  true  "Path where the image is located"
// @Param        imageName  path      string  true  "Name of the image to delete"
// @Success      200        {object}  map[string]string
// @Failure      400        {object}  error
// @Failure      500        {object}  error
// @Router       /api/v2/images/{path}/{imageName} [delete]
func (h *Handler) DeleteImage(c *gin.Context) {
	path := c.Param("path")
	imageName := c.Param("imageName")

	path, err := h.Services.ImageService.DeleteImage(c.Request.Context(), path, imageName)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, path)
}
