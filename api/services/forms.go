package services

import (
	"workerbee/models"
	"workerbee/repositories"
)

type FormService struct {
	repo repositories.Formrepositories
}

func NewFormService(repo repositories.Formrepositories) *FormService {
	return &FormService{repo: repo}
}

func (s *FormService) GetForms(search, limit, offset, orderBy, sort string) ([]models.FormWithTotalCount, error) {
	return s.repo.GetForms(search, limit, offset, orderBy, sort)
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

func (s *FormService) DeleteForm(id string) (models.Form, error) {
	return s.repo.DeleteForm(id)
}
