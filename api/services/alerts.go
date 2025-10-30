package services

import "workerbee/repositories"

type AlertService struct {
	repo repositories.AlertRepository
}

func NewAlertService(repo repositories.AlertRepository) *AlertService {
	return &AlertService{repo: repo}
}