package repositories

import (
	"context"
	"image"
	"mime/multipart"
	"strconv"
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
	CreateAlbum(ctx context.Context, body models.CreateAlbum) (models.CreateAlbum, error)
	UploadImagesToAlbum(ctx context.Context, id string, files []*multipart.FileHeader) error
	GetAlbum(ctx context.Context, id string) (models.AlbumWithImages, error)
	GetAlbums(ctx context.Context, orderBy, sort, search string, limit int, offset int) ([]models.AlbumsWithTotalCount, error)
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

func (ar *albumsRepository) CreateAlbum(ctx context.Context, body models.CreateAlbum) (models.CreateAlbum, error) {
	return db.AddOneRow(
		ar.db,
		"./db/albums/post_album.sql",
		body,
	)
}

func (ar *albumsRepository) UploadImagesToAlbum(ctx context.Context, id string, files []*multipart.FileHeader) error {
	path := string(id) + "/"

	for _, file := range files {
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

func (ar *albumsRepository) GetAlbum(ctx context.Context, id string) (models.AlbumWithImages, error) {
	album, err := db.ExecuteOneRow[models.AlbumWithImages](ar.db, "./db/albums/get_album.sql", id)
	if err != nil {
		return models.AlbumWithImages{}, err
	}

	path := id

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	prefix := internal.ALBUM_PATH + path

	var images []string
	paginator := s3.NewListObjectsV2Paginator(ar.DO, &s3.ListObjectsV2Input{
		Bucket: aws.String(ar.Bucket),
		Prefix: aws.String(prefix),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return models.AlbumWithImages{}, err
		}

		for _, obj := range page.Contents {
			if strings.HasSuffix(*obj.Key, "/") {
				continue
			}

			images = append(images, strings.TrimPrefix(*obj.Key, prefix))
		}
	}

	album.Images = images

	return album, nil
}

func (ar *albumsRepository) GetAlbums(ctx context.Context, orderBy, sort, search string, limit int, offset int) ([]models.AlbumsWithTotalCount, error) {
	albums, err := db.FetchAllElements[models.AlbumsWithTotalCount](
		ar.db,
		"./db/albums/get_albums.sql",
		orderBy, sort,
		limit, offset,
		search,
	)
	if err != nil {
		return nil, err
	}

	images := make(map[string][]string)
	neededAlbums := make(map[string]bool)
	for _, album := range albums {
		albumID := strconv.Itoa(album.ID)
		neededAlbums[albumID] = true
	}

	paginator := s3.NewListObjectsV2Paginator(ar.DO, &s3.ListObjectsV2Input{
		Bucket: aws.String(ar.Bucket),
		Prefix: aws.String(internal.ALBUM_PATH),
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

			parts := strings.Split(*obj.Key, "/")
			if len(parts) >= 2 {
				albumID := parts[1]

				if !neededAlbums[albumID] {
					continue
				}

				if len(images[albumID]) < 3 {
					filename := parts[len(parts)-1]
					images[albumID] = append(images[albumID], filename)

					if len(images[albumID]) >= 3 {
						delete(neededAlbums, albumID)
					}
				}
			}
		}
	}
	for i := range albums {
		albumID := strconv.Itoa(albums[i].ID)
		albums[i].Images = images[albumID]
	}

	return albums, nil
}
