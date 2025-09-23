package repositories

import (
	"workerbee/db"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type LocationRepository interface {
	GetLocations(search, limit, offset, orderBy, sort string) ([]models.LocationWithTotalCount, error)
}

type locationRepository struct {
	db *sqlx.DB
}

func NewLocationRepository(db *sqlx.DB) LocationRepository {
	return &locationRepository{db: db}
}

func (r *locationRepository) GetLocations(search, limit, offset, orderBy, sort string) ([]models.LocationWithTotalCount, error) {
	locations, err := db.FetchAllElements[models.LocationWithTotalCount](
		r.db,
		"./db/locations/get_locations.sql",
		orderBy, sort,
		limit,
		offset,
		search,
	)
	if err != nil {
		return nil, err
	}
	return locations, nil
}
