package repositories

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"image"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"golang.org/x/sync/semaphore"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/jmoiron/sqlx"
)

type AlbumsRepository interface {
	CreateAlbum(ctx context.Context, body models.CreateAlbum) (models.CreateAlbum, error)
	UploadImagesToAlbum(ctx context.Context, id string, files []*multipart.FileHeader) error
	GetAlbum(ctx context.Context, id string) (models.AlbumWithImages, error)
	GetAlbums(ctx context.Context, orderBy, sort, search string, limit int, offset int) ([]models.AlbumsWithTotalCount, error)
	UpdateAlbum(ctx context.Context, body models.CreateAlbum) (models.CreateAlbum, error)
	DeleteAlbum(ctx context.Context, id string) (int, error)
	DeleteAlbumImage(ctx context.Context, key, id string) error
	SetAlbumCover(ctx context.Context, id string, imageName string) error
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

	type uploadResult struct {
		err error
	}

	results := make(chan uploadResult, len(files))
	semaphore := semaphore.NewWeighted(20)

	for _, file := range files {
		if err := semaphore.Acquire(ctx, 1); err != nil {
			return err
		}

		go func(f *multipart.FileHeader) {
			defer semaphore.Release(1)
			src, err := f.Open()
			if err != nil {
				results <- uploadResult{err: err}
				return
			}
			defer src.Close()

			img, _, err := image.Decode(src)
			if err != nil {
				results <- uploadResult{err: err}
				return
			}

			width, height := img.Bounds().Dx(), img.Bounds().Dy()
			newWidth, newHeight := internal.DownscaleImage(width, height)

			if newWidth != width || newHeight != height {
				img = imaging.Resize(img, newWidth, newHeight, imaging.Box)
			}

			quality := 85
			var buf bytes.Buffer

			for quality >= 40 {
				buf.Reset()
				err = webp.Encode(&buf, img, &webp.Options{Lossless: false, Quality: float32(quality)})
				if err != nil {
					results <- uploadResult{err: err}
					return
				}

				if buf.Len() <= int(internal.MaxAlbumImageSize) {
					break
				}

				quality -= 10
			}

			for buf.Len() > int(internal.MaxAlbumImageSize) {
				bounds := img.Bounds()
				width := bounds.Dx()
				height := bounds.Dy()

				newWidth := int(float64(width) * .9)
				newHeight := int(float64(height) * .9)

				if newWidth < 100 || newHeight < 100 {
					break
				}

				img = imaging.Resize(img, newWidth, newHeight, imaging.Lanczos)

				buf.Reset()
				err = webp.Encode(&buf, img, &webp.Options{Lossless: false, Quality: 75})
				if err != nil {
					results <- uploadResult{err: err}
					return
				}
			}
			uploadPath := path
			if !strings.HasSuffix(uploadPath, "/") {
				uploadPath += "/"
			}

			filename := strings.TrimSuffix(f.Filename, filepath.Ext(f.Filename)) + ".webp"

			randomBytes := make([]byte, 6)
			_, err = rand.Read(randomBytes)
			if err != nil {
				results <- uploadResult{err: err}
				return
			}

			hash := sha256.Sum256(randomBytes)

			key := internal.ALBUM_PATH + uploadPath + "img_" + hex.EncodeToString(hash[:4]) + "_" + filename

			contentType := "image/webp"

			_, err = ar.DO.PutObject(ctx, &s3.PutObjectInput{
				Bucket:      aws.String(ar.Bucket),
				Key:         aws.String(key),
				Body:        bytes.NewReader(buf.Bytes()),
				ACL:         types.ObjectCannedACLPublicRead,
				ContentType: aws.String(contentType),
			})

			img = nil
			buf.Reset()

			results <- uploadResult{err: err}
		}(file)
	}
	var errors []error
	for range files {
		result := <-results
		if result.err != nil {
			errors = append(errors, result.err)
		}
	}
	close(results)

	if len(errors) > 0 {
		return errors[0]
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

func (ar *albumsRepository) UpdateAlbum(ctx context.Context, body models.CreateAlbum) (models.CreateAlbum, error) {
	return db.AddOneRow(
		ar.db,
		"./db/albums/put_album.sql",
		body,
	)
}

func (ar *albumsRepository) DeleteAlbum(ctx context.Context, id string) (int, error) {
	returnID, err := db.ExecuteOneRow[int](
		ar.db,
		"./db/albums/delete_album.sql",
		id,
	)

	var continuationToken *string

	prefix := internal.ALBUM_PATH + id
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	for {
		listOutput, err := ar.DO.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
			Bucket:            aws.String(ar.Bucket),
			Prefix:            aws.String(prefix),
			ContinuationToken: continuationToken,
		})
		if err != nil {
			return 0, err
		}

		if len(listOutput.Contents) == 0 {
			break
		}

		var objectsToDelete []types.ObjectIdentifier
		for _, obj := range listOutput.Contents {
			objectsToDelete = append(objectsToDelete, types.ObjectIdentifier{
				Key: obj.Key,
			})
		}

		_, err = ar.DO.DeleteObjects(ctx, &s3.DeleteObjectsInput{
			Bucket: aws.String(ar.Bucket),
			Delete: &types.Delete{
				Objects: objectsToDelete,
				Quiet:   aws.Bool(true),
			},
		})
		if err != nil {
			return 0, err
		}

		if *listOutput.IsTruncated {
			continuationToken = listOutput.NextContinuationToken
		} else {
			break
		}
	}
	return returnID, err
}

func (ar *albumsRepository) DeleteAlbumImage(ctx context.Context, key, id string) error {
	_, err := db.ExecuteOneRow[models.AlbumWithImages](
		ar.db,
		"./db/albums/get_album.sql",
		id,
	)
	if err != nil {
		return err
	}

	_, err = ar.DO.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(ar.Bucket),
		Key:    aws.String(key),
	})
	return err
}

func (ar *albumsRepository) SetAlbumCover(ctx context.Context, id string, imageName string) error {
	prefix := internal.ALBUM_PATH + id + "/"

	listOutput, err := ar.DO.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:  aws.String(ar.Bucket),
		Prefix:  aws.String(prefix + "coverimg_"),
		MaxKeys: aws.Int32(1),
	})
	if err != nil {
		return err
	}

	if len(listOutput.Contents) > 0 {
		firstKey := *listOutput.Contents[0].Key
		if strings.Contains(firstKey, "coverimg_") {
			oldCoverName := strings.Replace(firstKey, "coverimg_", "img_", 1)

			_, err = ar.DO.CopyObject(ctx, &s3.CopyObjectInput{
				Bucket:     aws.String(ar.Bucket),
				CopySource: aws.String(ar.Bucket + "/" + firstKey),
				Key:        aws.String(oldCoverName),
			})
			if err != nil {
				return err
			}

			_, err = ar.DO.DeleteObject(ctx, &s3.DeleteObjectInput{
				Bucket: aws.String(ar.Bucket),
				Key:    aws.String(firstKey),
			})
			if err != nil {
				return err
			}
		}
	}
	coverImageName := strings.Replace(imageName, "img_", "coverimg_", 1)
	path := internal.ALBUM_PATH + id + "/" + imageName
	coverPath := internal.ALBUM_PATH + id + "/" + coverImageName

	_, err = ar.DO.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(ar.Bucket),
		CopySource: aws.String(ar.Bucket + "/" + path),
		Key:        aws.String(coverPath),
	})
	if err != nil {
		return err
	}

	_, err = ar.DO.PutObjectAcl(ctx, &s3.PutObjectAclInput{
		Bucket: aws.String(ar.Bucket),
		Key:    aws.String(coverPath),
		ACL:    types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return err
	}

	_, err = ar.DO.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(ar.Bucket),
		Key:    aws.String(internal.ALBUM_PATH + id + "/" + imageName),
	})
	return err
}
