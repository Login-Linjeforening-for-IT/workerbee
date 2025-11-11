package services

import (
	"strconv"
	"strings"
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repositories"
)

var allowedSortColumnsAlerts = map[string]string{
	"id":             "a.id",
	"service":        "a.service",
	"page":           "a.page",
	"title_en":       "a.title_en",
	"title_no":       "a.title_no",
	"description_en": "a.description_en",
	"description_no": "a.description_no",
}

type AlertService struct {
	repo repositories.AlertRepository
}

func NewAlertService(repo repositories.AlertRepository) *AlertService {
	return &AlertService{repo: repo}
}

func (s *AlertService) CreateAlert(alert models.Alert) (models.Alert, error) {
	if !strings.HasPrefix(alert.Page, "/") {
		alert.Page = "/" + alert.Page
	}

	alert, err := s.repo.CreateAlert(alert)
	if err != nil {
		return models.Alert{}, err
	}



	return alert, nil
}

func (s *AlertService) GetAllAlerts(search, limit_str, offset_str, orderBy, sort string) ([]models.AlertWithTotalCount, error) {
	orderBySanitized, sortSanitized, err := internal.SanitizeSort(orderBy, sort, allowedSortColumnsAlerts)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetAllAlerts(limit, offset, search, orderBySanitized, sortSanitized)
}

func (s *AlertService) GetAlertByServiceAndPage(service, page string) (models.Alert, error) {
	if !strings.HasPrefix(page, "/") {
		page = "/" + page
	}

	return s.repo.GetAlertByServiceAndPage(service, page)
}

func (s *AlertService) GetAlertByID(id_str string) (models.Alert, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.Alert{}, internal.ErrInvalid
	}

	return s.repo.GetAlertByID(id)
}

func (s *AlertService) UpdateAlert(id_str string, alert models.Alert) (models.Alert, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.Alert{}, internal.ErrInvalid
	}
	alert.ID = id

	if !strings.HasPrefix(alert.Page, "/") {
		alert.Page = "/" + alert.Page
	}

	return s.repo.UpdateAlert(alert)
}

func (s *AlertService) DeleteAlert(id string) (int, error) {
	return s.repo.DeleteAlert(id)
}
