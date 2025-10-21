package services

import (
	"workerbee/repositories"
)

type HoneyService struct {
	repo repositories.HoneyRepository
}

func NewHoneyService(repo repositories.HoneyRepository) *HoneyService {
	return &HoneyService{repo: repo}
}

func (s *HoneyService) GetTextServices() ([]string, error) {
	return s.repo.GetTextServices()
}

func (s *HoneyService) GetAllPathsInService(service string) (map[string][]string, error) {
	rows, err := s.repo.GetAllPathsInService(service)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]string)
	for _, row := range rows {
		result[row.Page] = row.Languages
	}
	return result, nil
}
