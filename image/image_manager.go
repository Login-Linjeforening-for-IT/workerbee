package image

import (
	"context"
	"errors"
	"fmt"
	"image"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Requests interface {
	UploadImage(ctx context.Context, it imageType, file *os.File, fileName string)
	FetchImageList(ctx context.Context)
}

type Manager struct {
	do *DOStore.doStore
	*Store
}

type imageType int

const (
	ImageTypeJobs imageType = iota
	ImageTypeEventsBanner
	ImageTypeEventsSmall
	ImageTypeOrg
)

func (reqs *Requests) uploadToS3(localFilePath, fileName, doPath string) error {
	file, err := os.Open(localFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	object := s3.PutObjectInput{
		Bucket: aws.String("beehive"),
		Key:    aws.String(doPath + fileName),
		Body:   file,
		ACL:    aws.String("public-read"),
		Metadata: map[string]*string{
			"x-amz-meta-my-key": aws.String("your-value"),
		},
	}

	_, err = s3Client.PutObject(&object)
	return err
}

func (server *Server) UploadImage() error {
	// Create a temporary file to store the uploaded file
	tempFile, err := os.CreateTemp("", "uploaded-*.png")
	if err != nil {
		return errors.New("failed to create temporary file")
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Copy the file to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		return errors.New("failed to copy file")
	}

	// Check that the file follows our rules
	if err := checkUploadedImage(tempFile, ratioW, ratioH); err != nil {
		return fmt.Errorf("file check failed: %s", err.Error())
	}

	// Upload the file to S3
	err = server.uploadToS3(tempFile.Name(), fileName, folderPath)
	if err != nil {
		return errors.New("failed to upload file to S3")
	}

	return nil
}

type DropDownFileItem struct {
	Name     string `json:"name"`
	Size     string `json:"size"`
	Filepath string `json:"filepath"`
}

func (server *Server) FetchImageList(prefix string) ([]DropDownFileItem, error) {
	// List objects in the specified bucket with the given prefix
	input := s3.ListObjectsInput{
		Bucket: aws.String("beehive"),
		Prefix: aws.String(prefix),
	}

	result, err := s3Client.ListObjects(&input)
	if err != nil {
		return nil, fmt.Errorf("failed to list objects: %s", err.Error())
	}

	var images []DropDownFileItem

	for _, object := range result.Contents {
		// Extract file information from the object metadata
		name := removePrefix(*object.Key, prefix)

		// Skip items representing the folder itself
		if name == "" {
			continue
		}

		// Check if the Size field is not nil
		var size string
		if object.Size != nil {
			size = byteConverter(*object.Size, 2)
		} else {
			size = "0 B" // Default value if Size is nil
		}

		filepath := *object.Key

		images = append(images, DropDownFileItem{name, size, filepath})
	}

	return images, nil
}

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

func checkFileType(file *os.File) (string, error) {
	if _, err := file.Seek(0, 0); err != nil {
		return "", errors.New("failed to seek to the beginning of the file")
	}

	// Detect Content-Type using net/http
	buffer := make([]byte, 512) // Read the first 512 bytes to detect Content-Type
	_, err := file.Read(buffer)
	if err != nil {
		return "", errors.New("failed to read file for Content-Type detection")
	}

	contentType := http.DetectContentType(buffer)
	// Check if the contentType is of the allowed formats
	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/gif" {
		return "", errors.New("invalid image type")
	}

	return contentType, nil
}

func checkFileRatio(file *os.File, ratioW, ratioH int) error {
	// Seek to the beginning of the file
	_, err := file.Seek(0, 0)
	if err != nil {
		return err
	}

	// Open the file
	imageFile, _, err := image.Decode(file)
	if err != nil {
		return errors.New("failed to decode image")
	}

	// Get image dimensions
	bounds := imageFile.Bounds()
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

func checkFileSize(file *os.File, fileType string) error {
	// Get file information
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// Check file size
	sizeBytes := fileInfo.Size()
	switch fileType {
	case "image/jpeg", "image/jpg":
		if sizeBytes > 500*1024 {
			return errors.New("file size exceeds maximum allowed size")
		}
	case "image/png":
		if sizeBytes > 500*1024 {
			return errors.New("file size exceeds maximum allowed size")
		}
	case "image/gif":
		if sizeBytes > 2000*1024 {
			return errors.New("file size exceeds maximum allowed size")
		}
	default:
		if sizeBytes > 500*1024 {
			return errors.New("file size exceeds maximum allowed size")
		}
	}

	return nil
}

func checkUploadedImage(file *os.File, ratioW, ratioH int) error {
	fileType, err := checkFileType(file)
	if err != nil {
		return err
	}

	err = checkFileSize(file, fileType)
	if err != nil {
		return err
	}

	err = checkFileRatio(file, ratioW, ratioH)
	if err != nil {
		return err
	}

	return nil
}
