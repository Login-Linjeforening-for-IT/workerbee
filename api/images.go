package api

import (
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

func (server *Server) uploadToS3(localFilePath, fileName, doPath string) error {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(server.config.DOKey, server.config.DOSecret, ""),
		Endpoint:         aws.String("https://ams3.digitaloceanspaces.com"),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String("ams3"),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)

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

func checkFileType(file *os.File) (string, error) {
	// Check image type
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
	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/gif" {
		return "", errors.New("invalid image type")
	}

	return contentType, nil
}

func checkFileRatio(file *os.File, ratioW, ratioH int) error {
	// Print or log information about the file
	fmt.Println("Checking file for ratio:", file.Name())

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

func (server *Server) uploadImageRequest(ctx *gin.Context, folderPath string, ratioW, ratioH int) {
	// Get the file from the request
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// Create a temporary file to store the uploaded file
	tempFile, err := os.CreateTemp("", "uploaded-*.png")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temporary file"})
		return
	}

	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Copy the file to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy file"})
		return
	}

	// Check that the file follows our rules
	if err := checkUploadedImage(tempFile, ratioW, ratioH); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("File check failed: %s", err.Error())})
		return
	}

	// Upload the file to S3
	err = server.uploadToS3(tempFile.Name(), header.Filename, folderPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to S3"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func (server *Server) uploadEventImageBanner(ctx *gin.Context) {
	server.uploadImageRequest(ctx, "img/events/banner/", 10, 4)
}

func (server *Server) uploadEventImageSmall(ctx *gin.Context) {
	server.uploadImageRequest(ctx, "img/events/small/", 10, 4)
}

func (server *Server) uploadAdImage(ctx *gin.Context) {
	server.uploadImageRequest(ctx, "img/ads/", 3, 2)
}

func (server *Server) uploadOrganizationImage(ctx *gin.Context) {
	server.uploadImageRequest(ctx, "img/organizations/", 3, 2)
}
