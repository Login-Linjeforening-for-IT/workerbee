package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"slices"
	"strings"
	"workerbee/config"
	"workerbee/internal"

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

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	key := "img/" + path + file.Filename

	_, err = is.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(is.Bucket),
		Key:         aws.String(key),
		Body:        src,
		ACL:         types.ObjectCannedACLPublicRead,
		ContentType: aws.String(file.Header.Get("Content-Type")),
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s", internal.CDN_URL, key), nil
}
