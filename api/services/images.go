package services

import (
	"context"
	"image"
	"mime/multipart"
	"slices"
	"strings"
	"workerbee/internal"
	"workerbee/repositories"

	_ "image/jpeg"
)

var validPaths = []string{
	"events",
	"jobs",
	"organizations",
}

var maxImageSizeMB = int64(1000000)
var imageRatio = 2.5

type ImageService struct {
	repo repositories.ImageRepository
}

func NewImageService(repo repositories.ImageRepository) *ImageService {
	return &ImageService{
		repo: repo,
	}
}

func (is *ImageService) UploadImage(file *multipart.FileHeader, ctx context.Context, path string) (string, error) {
	if !slices.Contains(validPaths, path) {
		return "", internal.ErrInvalidImagePath
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	img, format, err := image.Decode(src)
	if err != nil {
		return "", err
	}

	src.Seek(0, 0)

	contentType := "image/" + format

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	ratio := float64(width) / float64(height)
	if ratio != imageRatio {
		return "", internal.ErrInvalidImageRatio
	}

	if file.Size > maxImageSizeMB {
		return "", internal.ErrImageTooLarge
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	key := internal.IMG_PATH + path + file.Filename

	err = is.repo.UploadImage(ctx, key, contentType, src)
	if err != nil {
		return "", err
	}
	return file.Filename, nil
}

func (is *ImageService) GetImagesInPath(ctx context.Context, path string) ([]string, error) {
	if !slices.Contains(validPaths, path) {
		return nil, internal.ErrInvalidImagePath
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	prefix := internal.IMG_PATH + path

	images, err := is.repo.GetImagesInPath(ctx, prefix)
	if err != nil {
		return nil, err
	}

	return images, nil
}

func (is *ImageService) DeleteImage(ctx context.Context, path, imageName string) (string, error) {
	if !slices.Contains(validPaths, path) {
		return "", internal.ErrInvalidImagePath
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	key := internal.IMG_PATH + path + imageName

	err := is.repo.DeleteImage(ctx, key)
	if err != nil {
		return "", err
	}

	return key, nil
}
