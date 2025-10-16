package repositories

import (
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type LocationRepository interface {
	CreateLocation(location models.NewLocation) (models.NewLocation, error)
	GetLocations(limit, offset int, search, orderBy, sort string, types []string) ([]models.LocationWithTotalCount, error)
	GetLocation(id string) (models.Location, error)
	UpdateLocation(location models.NewLocation) (models.NewLocation, error)
	DeleteLocation(id string) (int, error)
	GetAllLocationTypes() (string, error)
}

type locationRepository struct {
	db *sqlx.DB
}

func NewLocationRepository(db *sqlx.DB) LocationRepository {
	return &locationRepository{db: db}
}

func (r *locationRepository) CreateLocation(location models.NewLocation) (models.NewLocation, error) {
	return db.AddOneRow(
		r.db,
		"./db/locations/post_location.sql",
		location,
	)
}

func (r *locationRepository) UpdateLocation(location models.NewLocation) (models.NewLocation, error) {
	return db.AddOneRow(
		r.db,
		"./db/locations/put_location.sql",
		location,
	)
}

func (r *locationRepository) GetLocations(limit, offset int, search, orderBy, sort string, types []string) ([]models.LocationWithTotalCount, error) {
	locations, err := db.FetchAllElements[models.LocationWithTotalCount](
		r.db,
		"./db/locations/get_locations.sql",
		orderBy, sort,
		limit,
		offset,
		search,
		pq.Array(types),
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

func (r *locationRepository) GetAllLocationTypes() (string, error) {
	locationTypes, err := db.FetchAllEnumTypes(
		r.db,
		"./db/locations/get_all_location_types.sql",
	)
	if err != nil {
		return "", err
	}
	return locationTypes, nil
}

func (r *locationRepository) DeleteLocation(id string) (int, error) {
	locationId, err := db.ExecuteOneRow[int](
		r.db, "./db/locations/delete_location.sql", id,
	)
	if err != nil {
		return 0, internal.ErrInvalid
	}
	return locationId, nil
}
