package services

import (
	"strconv"
	"strings"
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repositories"
)

var allowedSortColumnsLocs = map[string]string{
	"id":         "l.id",
	"name_no":    "l.name_no",
	"name_en":    "l.name_en",
	"type":       "l.type",
	"city_name":  "city_name",
	"created_at": "l.created_at",
	"updated_at": "l.updated_at",
}

type LocationService struct {
	repo repositories.LocationRepository
}

func NewLocationService(repo repositories.LocationRepository) *LocationService {
	return &LocationService{repo: repo}
}

func (s *LocationService) CreateLocation(location models.Location) (models.Location, error) {
	return s.repo.CreateLocation(location)
}

func (s *LocationService) GetLocations(search, limit, offset, orderBy, sort, types string) ([]models.LocationWithTotalCount, error) {
	var err error
	
	orderBySanitized, sortSanitized, ok := internal.SanitizeSort(orderBy, sort, allowedSortColumnsLocs)
	if ok != nil {
		return nil, internal.ErrInvalid
	}

	var typeSlice []string
	if types != "" {
		typeSlice, err = internal.ParseCSVToSlice[string](types)
		if err != nil {
			return nil, err
		}
	} else {
		typeSlice = make([]string, 0)
	}

	return s.repo.GetLocations(search, limit, offset, orderBySanitized, strings.ToUpper(sortSanitized), typeSlice)
}

func (s *LocationService) GetLocation(id string) (models.Location, error) {
	return s.repo.GetLocation(id)
}

func (s *LocationService) DeleteLocation(id string) (models.Location, error) {
	return s.repo.DeleteLocation(id)
}

func (s *LocationService) UpdateLocation(id_str string, location models.Location) (models.Location, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.Location{}, err
	}

	location.ID = id

	return s.repo.UpdateLocation(location)
}
