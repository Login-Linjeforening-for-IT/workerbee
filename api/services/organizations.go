package services

import "workerbee/repositories"

type OrganizationService struct {
	repo repositories.OrganizationRepository
}

func NewOrganizationService(repo repositories.OrganizationRepository) *OrganizationService {
	return &OrganizationService{repo: repo}
}
