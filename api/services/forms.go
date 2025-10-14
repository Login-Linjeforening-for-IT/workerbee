package services

import (
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repositories"
)

type FormService struct {
	repo repositories.Formrepositories
}

func NewFormService(repo repositories.Formrepositories) *FormService {
	return &FormService{repo: repo}
}

func (s *FormService) GetForms(search, limit_str, offset_str, orderBy, sort string) ([]models.FormWithTotalCount, error) {
	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetForms(limit, offset, search, orderBy, sort)
}

func (s *FormService) GetForm(id string) (*models.FormWithQuestion, error) {
	return s.repo.GetForm(id)
}

func (s *FormService) PostForm(form models.Form) (models.Form, error) {
	return s.repo.PostForm(form)
}

func (s *FormService) PutForm(id string, form models.Form) (models.Form, error) {
	return s.repo.PutForm(id, form)
}

func (s *FormService) DeleteForm(id string) (int, error) {
	return s.repo.DeleteForm(id)
}
