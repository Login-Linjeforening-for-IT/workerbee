package services

import "workerbee/repositories"

type LocationService struct {
	repo repositories.LocationRepository
}

func NewLocationService(repo repositories.LocationRepository) *LocationService {
	return &LocationService{repo: repo}
}
