package repositories

import (
	"context"
	"image"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"workerbee/internal"
	"workerbee/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/jmoiron/sqlx"
)

var maxAlbumImageSize = int64(2000000)

type AlbumsRepository interface {
	CreateAlbum(ctx context.Context, images []*multipart.FileHeader, body models.Album) error
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

func (ar *albumsRepository) CreateAlbum(ctx context.Context, images []*multipart.FileHeader, body models.Album) error {
	tx, err := ar.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	sqlBytes, err := os.ReadFile("./db/albums/create_album.sql")
	if err != nil {
		return err
	}

	row := tx.QueryRow(
		string(sqlBytes),
		body.NameEn,
		body.NameNo,
		body.DescriptionEn,
		body.DescriptionNo,
	)
	var newAlbumID int
	err = row.Scan(&newAlbumID)
	if err != nil {
		return err
	}

	path := strconv.Itoa(newAlbumID) + "/"

	for _, file := range images {
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

	return tx.Commit()
}
