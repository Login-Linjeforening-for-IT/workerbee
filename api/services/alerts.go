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