package services

import (
	"workerbee/config"
	"workerbee/internal"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

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
