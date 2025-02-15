/*
Package images provides a simple interface to store and retrieve a list of images from a backend.
The DigitalOcean store implementation is used to store the images on DigitalOcean Spaces.
*/

package images

import (
	"context"
	"mime"
	"path/filepath"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type DOStore struct {
	client *s3.Client
	bucket string
}

var _ Store = &DOStore{}

type DOConfig struct {
	DOKey     string `config:"DO_ACCESS_KEY_ID"`
	DOSecret  string `config:"DO_SECRET_ACCESS_KEY"`
	DORegion  string `config:"DO_REGION" default:"ams3"`
	DOBaseURL string `config:"DO_BASE_URL" default:"https://ams3.digitaloceanspaces.com"`
	DOBucket  string `config:"DO_BUCKET" default:"beehive"`
}

// Authenticates wit DO and returns a client
func NewDOStore(
	doConfig *DOConfig,
) (*DOStore, error) {
	creds := credentials.NewStaticCredentialsProvider(doConfig.DOKey, doConfig.DOSecret, "")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(doConfig.DORegion),
		config.WithCredentialsProvider(creds),
	)

	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(doConfig.DOBaseURL)
	})

	return &DOStore{
		client: client,
		bucket: doConfig.DOBucket,
	}, nil
}

func (store *DOStore) GetImages(dir string) ([]FileDetails, error) {
	result, err := store.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(store.bucket),
		Prefix: aws.String(dir),
	})
	if err != nil {
		return nil, err
	}

	pathRegex := regexp.MustCompile(`^(.*/)`)
	fileRegex := regexp.MustCompile(`([^/]+)$`)

	var files []FileDetails
	for _, obj := range result.Contents {
		fileName := fileRegex.FindString(*obj.Key)
		// Result contains the directory itself, skip it
		if fileName == "" {
			continue
		}

		files = append(files, FileDetails{
			Name: fileName,
			Size: *obj.Size,
			Path: pathRegex.FindString(*obj.Key),
		})
	}

	return files, nil
}

func (store *DOStore) UploadImage(dir string, fileName string, file File) error {
	mimeType := mime.TypeByExtension(filepath.Ext(fileName))
	// Do we want to include the id/name of the one who uploaded the file? (x-amz-meta-)
	_, err := store.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(store.bucket),
		Key:         aws.String(dir + fileName),
		Body:        file,
		ACL:         types.ObjectCannedACLPublicRead,
		ContentType: &mimeType,
	})
	if err != nil {
		return err
	}

	return nil
}
