package repositories

import "github.com/jmoiron/sqlx"

type LocationRepository interface {
}

type locationRepository struct {
	db *sqlx.DB
}

func NewLocationRepository(db *sqlx.DB) LocationRepository {
	return &locationRepository{db: db}
}
