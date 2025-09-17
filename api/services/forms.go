package services

import (
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/repository"
)

type FormService struct {
	repo repository.FormRepository
}

func NewFormService(repo repository.FormRepository) *FormService {
	return &FormService{repo: repo}
}

func (s *FormService) GetForms(search, limit, offset, orderBy, sort string) ([]models.FormWithTotalCount, error) {
	return s.repo.GetForms(search, limit, offset, orderBy, sort)
}

func (s *FormService) GetForm(id string) ([]models.Form, error) {
	return s.repo.GetForm(id)
}