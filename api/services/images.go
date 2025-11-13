package services

import (
	"bytes"
	"context"
	"image"
	"mime/multipart"
	"slices"
	"strings"
	"workerbee/internal"
	"workerbee/repositories"

	"image/jpeg"
	"image/png"

	"github.com/chai2010/webp"
)

var validPaths = []string{
	"events",
	"jobs",
	"organizations",
}

var maxImageSize = int64(1000000)
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

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	if file.Size > maxImageSize {
		return "", internal.ErrImageTooLarge
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	var img image.Image
	filename := strings.ToLower(file.Filename)
	if strings.HasSuffix(filename, ".png") {
		img, err = png.Decode(src)
	} else if strings.HasSuffix(filename, ".jpg") || strings.HasSuffix(filename, ".jpeg") {
		img, err = jpeg.Decode(src)
	} else {
		return "", internal.ErrUnknownImageFormat
	}
	if err != nil {
		return "", err
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	ratio := float64(width) / float64(height)
	if ratio > imageRatio {
		return "", internal.ErrInvalidImageRatio
	}

	buf := new(bytes.Buffer)
	err = webp.Encode(buf, img, &webp.Options{Lossless: false, Quality: 80})
	if err != nil {
		return "", err
	}

	key := internal.IMG_PATH + path + strings.Split(file.Filename, ".")[0] + ".webp"

	err = is.repo.UploadImage(ctx, key, "image/webp", buf)
	if err != nil {
		return "", err
	}

	return key, nil
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
