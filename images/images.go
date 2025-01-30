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

// TODO: Implement use of constants
type ImageRatio struct {
	width  int
	height int
}

var Banner = ImageRatio{10, 4}
var Small = ImageRatio{10, 4}
var Ads = ImageRatio{3, 2}

// func (server *Server) uploadToS3(localFilePath, fileName, doPath string) error {
// 	file, err := os.Open(localFilePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	object := s3.PutObjectInput{
// 		Bucket: aws.String("beehive"),
// 		Key:    aws.String(doPath + fileName),
// 		Body:   file,
// 		ACL:    aws.String("public-read"),
// 		Metadata: map[string]*string{
// 			"x-amz-meta-my-key": aws.String("your-value"),
// 		},
// 	}

// 	_, err = s3Client.PutObject(&object)
// 	return err
// }

// func (server *Server) UploadImageRequest(ctx *gin.Context, folderPath string, ratioW, ratioH int) {
// 	// Get the file from the request
// 	file, header, err := ctx.Request.FormFile("file")
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer file.Close()

// 	// Create a temporary file to store the uploaded file
// 	tempFile, err := os.CreateTemp("", "uploaded-*.png")
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temporary file"})
// 		return
// 	}

// 	defer os.Remove(tempFile.Name())
// 	defer tempFile.Close()

// 	// Copy the file to the temporary file
// 	_, err = io.Copy(tempFile, file)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy file"})
// 		return
// 	}

// 	// Check that the file follows our rules
// 	if err := checkUploadedImage(tempFile, ratioW, ratioH); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("File check failed: %s", err.Error())})
// 		return
// 	}

// 	// Upload the file to S3
// 	err = server.uploadToS3(tempFile.Name(), header.Filename, folderPath)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to S3"})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
// }

// type DropDownFileItem struct {
// 	Name     string `json:"name"`
// 	Size     string `json:"size"`
// 	Filepath string `json:"filepath"`
// }

// func (server *Server) FetchImageList(ctx *gin.Context, prefix string) {
// 	// List objects in the specified bucket with the given prefix
// 	input := s3.ListObjectsInput{
// 		Bucket: aws.String("beehive"),
// 		Prefix: aws.String(prefix),
// 	}

// 	result, err := s3Client.ListObjects(&input)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to list objects: %s", err.Error())})
// 		return
// 	}

// 	var images []DropDownFileItem

// 	for _, object := range result.Contents {
// 		// Extract file information from the object metadata
// 		name := removePrefix(*object.Key, prefix)

// 		// Skip items representing the folder itself
// 		if name == "" {
// 			continue
// 		}

// 		// Check if the Size field is not nil
// 		var size string
// 		if object.Size != nil {
// 			size = byteConverter(*object.Size, 2)
// 		} else {
// 			size = "0 B" // Default value if Size is nil
// 		}

// 		filepath := *object.Key

// 		images = append(images, DropDownFileItem{name, size, filepath})
// 	}

// 	ctx.JSON(http.StatusOK, images)
// }

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
		if size > 500*1024 {
			return errors.New("file size exceeds maximum allowed size")
		}
	case "png":
		if size > 500*1024 {
			return errors.New("file size exceeds maximum allowed size")
		}
	case "gif":
		if size > 2000*1024 {
			return errors.New("file size exceeds maximum allowed size")
		}
	default:
		if size > 500*1024 {
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
