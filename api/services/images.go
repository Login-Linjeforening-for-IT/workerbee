package services

import (
	"context"
	"image"
	"math"
	"mime/multipart"
	"slices"
	"strings"
	"workerbee/config"
	"workerbee/internal"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var validPaths = []string{
	"events",
	"jobs",
	"organizations",
}

var maxImageSizeMB = int64(1000000)
var imageRatio = 2.5

type ImageService struct {
	Client *s3.Client
	Bucket string
}

func NewImageService() *ImageService {
	client := s3.New(s3.Options{
		Region: internal.REGION,
		Credentials: aws.NewCredentialsCache(
			credentials.StaticCredentialsProvider{
				Value: aws.Credentials{
					AccessKeyID:     config.DO_access_key_id,
					SecretAccessKey: config.DO_secret_access_key,
				},
			},
		),
		EndpointResolver: s3.EndpointResolverFromURL(config.DO_URL),
		UsePathStyle:     true,
	})

	return &ImageService{
		Client: client,
		Bucket: internal.BUCKET_NAME,
	}
}

func (is *ImageService) UploadImage(file *multipart.FileHeader, ctx context.Context, path string) (string, error) {
	if !slices.Contains(validPaths, path) {
		return "", internal.ErrInvalidImagePath
	}

	if is.Client == nil {
		return "", internal.ErrS3ClientNotInitialized
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
	if math.Abs(ratio-imageRatio) > 0.01 {
		return "", internal.ErrInvalidImageRatio
	}

	if file.Size > maxImageSizeMB {
		return "", internal.ErrImageTooLarge
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	key := internal.IMG_PATH + path + file.Filename

	_, err = is.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(is.Bucket),
		Key:         aws.String(key),
		Body:        src,
		ACL:         types.ObjectCannedACLPublicRead,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}
	// TODO futureproof this for folders
	return file.Filename, nil
}

func (is *ImageService) GetImagesInPath(ctx context.Context, path string) ([]string, error) {
	if !slices.Contains(validPaths, path) {
		return nil, internal.ErrInvalidImagePath
	}

	if is.Client == nil {
		return nil, internal.ErrS3ClientNotInitialized
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	prefix := internal.IMG_PATH + path
	var images []string

	paginator := s3.NewListObjectsV2Paginator(is.Client, &s3.ListObjectsV2Input{
		Bucket: aws.String(is.Bucket),
		Prefix: aws.String(prefix),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, obj := range page.Contents {
			if strings.HasSuffix(*obj.Key, "/") {
				continue
			}

			images = append(images, strings.TrimPrefix(*obj.Key, prefix))
		}
	}
	return images, nil
}

func (is *ImageService) DeleteImage(ctx context.Context, path, imageName string) (string, error) {
	if !slices.Contains(validPaths, path) {
		return "", internal.ErrInvalidImagePath
	}

	if is.Client == nil {
		return "", internal.ErrS3ClientNotInitialized
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	key := internal.IMG_PATH + path + imageName

	_, err := is.Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(is.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}

	return key, nil
}
