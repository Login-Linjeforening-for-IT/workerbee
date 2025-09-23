package repositories

import (
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type LocationRepository interface {
	GetLocations(search, limit, offset, orderBy, sort string) ([]models.LocationWithTotalCount, error)
	GetLocation(id string) (models.Location, error)
	DeleteLocation(id string) (models.Location, error)
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

func (r *locationRepository) GetLocation(id string) (models.Location, error) {
	location, err := db.ExecuteOneRow[models.Location](
		r.db, "./db/locations/get_location.sql", id,
	)
	if err != nil {
		return models.Location{}, internal.ErrInvalid
	}
	return location, nil
}

func (r *locationRepository) DeleteLocation(id string) (models.Location, error) {
	location, err := db.ExecuteOneRow[models.Location](
		r.db, "./db/locations/delete_location.sql", id,
	)
	if err != nil {
		return models.Location{}, internal.ErrInvalid
	}
	return location, nil
}
