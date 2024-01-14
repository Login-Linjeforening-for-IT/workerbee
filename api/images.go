package api

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func uploadToS3(localFilePath, fileName, doPath string) error {
	if err := godotenv.Load(); err != nil {
		fmt.Print("No .env file found")
	}

	key := os.Getenv("ACCESS_KEY_ID")
	secret := os.Getenv("SECRET_ACCESS_KEY")

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(key, secret, ""),
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

func (server *Server) uploadImageRequest(ctx *gin.Context, folderPath string) {
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

	// Upload the file to S3
	err = uploadToS3(tempFile.Name(), header.Filename, folderPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to S3"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func (server *Server) uploadEventImageBanner(ctx *gin.Context) {
	server.uploadImageRequest(ctx, "img/events/banner/")
}

func (server *Server) uploadEventImageSmall(ctx *gin.Context) {
	server.uploadImageRequest(ctx, "img/events/small/")
}

func (server *Server) uploadAdImage(ctx *gin.Context) {
	server.uploadImageRequest(ctx, "img/ads/")
}

func (server *Server) uploadOrganizationImage(ctx *gin.Context) {
	server.uploadImageRequest(ctx, "img/organizations/")
}
