package repositories

import (
	"context"
	"io"
	"strings"
	"workerbee/internal"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/jmoiron/sqlx"
)

type ImageRepository interface {
	UploadImage(ctx context.Context, key, contentType string, src io.Reader) error
	GetImagesInPath(ctx context.Context, prefix string) ([]string, error)
	DeleteImage(ctx context.Context, key string) error
}

type imageRepository struct {
	db     *sqlx.DB
	DO     *s3.Client
	Bucket string
}

func NewImageRepository(db *sqlx.DB, do *s3.Client) ImageRepository {
	return &imageRepository{
		db:     db,
		DO:     do,
		Bucket: internal.BUCKET_NAME,
	}
}

func (ir *imageRepository) UploadImage(ctx context.Context, key, contentType string, src io.Reader) error {
	_, err := ir.DO.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(ir.Bucket),
		Key:         aws.String(key),
		Body:        src,
		ACL:         types.ObjectCannedACLPublicRead,
		ContentType: aws.String(contentType),
	})
	return err
}

func (ir *imageRepository) GetImagesInPath(ctx context.Context, prefix string) ([]string, error) {
	var images []string
	paginator := s3.NewListObjectsV2Paginator(ir.DO, &s3.ListObjectsV2Input{
		Bucket: aws.String(ir.Bucket),
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

func (ir *imageRepository) DeleteImage(ctx context.Context, key string) error {
	_, err := ir.DO.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(ir.Bucket),
		Key:    aws.String(key),
	})
	return err
}
