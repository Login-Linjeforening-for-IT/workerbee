package repositories

import (
	"context"
	"image"
	"log"
	"mime/multipart"
	"strings"
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/jmoiron/sqlx"
)

var maxAlbumImageSize = int64(2000000)

type AlbumsRepository interface {
	CreateAlbum(ctx context.Context, body models.Album) (models.Album, error)
	UploadImagesToAlbum(ctx context.Context, id string, files []*multipart.FileHeader) error
}

type albumsRepository struct {
	db     *sqlx.DB
	DO     *s3.Client
	Bucket string
}

func NewAlbumsRepository(db *sqlx.DB, do *s3.Client) AlbumsRepository {
	return &albumsRepository{
		db:     db,
		DO:     do,
		Bucket: internal.BUCKET_NAME,
	}
}

func (ar *albumsRepository) CreateAlbum(ctx context.Context, body models.Album) (models.Album, error) {
	return db.AddOneRow(
		ar.db,
		"./db/albums/post_album.sql",
		body,
	)
}

func (ar *albumsRepository) UploadImagesToAlbum(ctx context.Context, id string, files []*multipart.FileHeader) error {
	path := string(id) + "/"

	for _, file := range files {
		log.Println("Uploading image:", file.Filename)

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		_, format, err := image.Decode(src)
		if err != nil {
			return err
		}

		_, err = src.Seek(0, 0)
		if err != nil {
			return err
		}

		if file.Size > maxAlbumImageSize {
			return internal.ErrImageTooLarge
		}

		if !strings.HasSuffix(path, "/") {
			path += "/"
		}

		key := internal.ALBUM_PATH + path + file.Filename

		contentType := "image/" + format

		_, err = ar.DO.PutObject(ctx, &s3.PutObjectInput{
			Bucket:      aws.String(ar.Bucket),
			Key:         aws.String(key),
			Body:        src,
			ACL:         types.ObjectCannedACLPublicRead,
			ContentType: aws.String(contentType),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
