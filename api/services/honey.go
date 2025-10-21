package services

import "workerbee/repositories"

type HoneyService struct {
	repo repositories.HoneyRepository
}

func NewHoneyService(repo repositories.HoneyRepository) *HoneyService {
	return &HoneyService{repo: repo}
}

func (s *HoneyService) GetTextServices() ([]string, error) {
	return s.repo.GetTextServices()
}
