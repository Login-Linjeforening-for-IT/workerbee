package repositories

import (
	"workerbee/internal"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jmoiron/sqlx"
)

type AlbumsRepository interface {
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
