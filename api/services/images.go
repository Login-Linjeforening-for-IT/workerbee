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
	"github.com/disintegration/imaging"
)

var validPaths = []string{
	"events",
	"jobs",
	"organizations",
}

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

	width, height := img.Bounds().Dx(), img.Bounds().Dy()

	newWidth, newHeight := internal.DownscaleImage(width, height)

	if newWidth != width || newHeight != height {
		img = imaging.Resize(img, newWidth, newHeight, imaging.Lanczos)
	}

	quality := 85
	var buf bytes.Buffer

	for quality >= 50 {
		buf.Reset()
		err = webp.Encode(&buf, img, &webp.Options{Lossless: false, Quality: float32(quality)})
		if err != nil {
			return "", err
		}

		if buf.Len() <= int(internal.MaxImageSize) {
			break
		}
		quality -= 10
	}

	for buf.Len() > int(internal.MaxImageSize) {
		bounds := img.Bounds()
		width := bounds.Dx()
		height := bounds.Dy()

		newWidth := int(float64(width) * .9)
		newHeight := int(float64(height) * .9)

		if newWidth < 100 || newHeight < 100 {
			break
		}

		img = imaging.Resize(img, newWidth, newHeight, imaging.Lanczos)

		buf.Reset()
		err = webp.Encode(&buf, img, &webp.Options{Lossless: false, Quality: 75})
		if err != nil {
			return "", err
		}
	}

	key := internal.IMG_PATH + path + strings.Split(file.Filename, ".")[0] + ".webp"

	err = is.repo.UploadImage(ctx, key, "image/webp", &buf)
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
