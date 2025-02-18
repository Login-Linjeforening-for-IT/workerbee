package images

import (
	"errors"
	"fmt"
	"image"
	"strings"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

const (
	BannerW = 10
	BannerH = 4
	SmallW  = 10
	SmallH  = 4
	AdsW    = 3
	AdsH    = 2
	OrgW    = 3
	OrgH    = 2
	MAX_SIZE = 5
	MB = 1048576
)

// Removes the prefix from the full path to the file. F.ex the prefix img/events/small from the full filepath
// img/events/small/testfile.png, and returns just the filename.
func removePrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

func byteConverter(size int64, precision int) string {
	const unit = 1024
	if size > unit*unit {
		return fmt.Sprintf("%.*f MiB", precision, float64(size)/(float64(unit)*float64(unit)))
	}
	return fmt.Sprintf("%.*f KiB", precision, float64(size)/float64(unit))
}

func checkFileRatio(img image.Image, ratioW, ratioH int) error {
	// Get image dimensions
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Check image ratio (example: allow images with w:h ratio)
	allowedRatio := float64(ratioW) / float64(ratioH)
	actualRatio := float64(width) / float64(height)

	// Check that the image has the correct aspect ratio, within some margin
	margin := allowedRatio * 0.05
	if !(actualRatio < (allowedRatio+margin) && actualRatio > (allowedRatio-margin)) {
		return errors.New("invalid image ratio")
	}

	return nil
}

// size is in bytes
// TODO: Update max file sizes?
func checkFileSize(size int64, fileType string) error {
	switch fileType {
	case "jpeg", "jpg":
		if size > MAX_SIZE * MB {
			return errors.New("file size exceeds maximum allowed size")
		}
	case "png":
		if size > MAX_SIZE * MB {
			return errors.New("file size exceeds maximum allowed size")
		}
	case "gif":
		if size > MAX_SIZE * MB {
			return errors.New("file size exceeds maximum allowed size")
		}
	default:
		if size > MAX_SIZE * MB {
			return errors.New("file size exceeds maximum allowed size")
		}
	}

	return nil
}

func CheckImage(file File, size int64, ratioW, ratioH int) error {
	img, imgType, err := image.Decode(file)
	if err != nil {
		return err
	}
	file.Seek(0, 0)

	err = checkFileRatio(img, ratioW, ratioH)
	if err != nil {
		return err
	}

	err = checkFileSize(size, imgType)
	if err != nil {
		return err
	}

	return checkFileRatio(img, ratioW, ratioH)
}
