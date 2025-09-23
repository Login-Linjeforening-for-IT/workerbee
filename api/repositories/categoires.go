package repositories

import "github.com/jmoiron/sqlx"

type CategoireRepository interface {
}

type categoireRepository struct {
	db *sqlx.DB
}

func NewCategoireRepository(db *sqlx.DB) CategoireRepository {
	return &categoireRepository{db: db}
}
