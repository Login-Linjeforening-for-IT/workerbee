package services

import "workerbee/repositories"

type AlertService struct {
	repo repositories.AlertRepository
}

func NewAlertService(repo repositories.AlertRepository) *AlertService {
	return &AlertService{repo: repo}
}

func (s *AlertService) GetAlertServices() ([]string, error) {
	return s.repo.GetAlertServices()
}

func (s *AlertService) GetAllPathsInAlertService(service string) (map[string][]string, error) {
	rows, err := s.repo.GetAllPathsInAlertService(service)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]string)
	for _, row := range rows {
		result[row.Page] = row.Languages
	}
	return result, nil
}
