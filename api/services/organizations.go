package services

import (
	"strings"
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repositories"
)

var allowedSortColumnsOrgs = map[string]string{
	"id":         "o.id",
	"name_no":    "o.title_no",
	"name_en":    "o.title_en",
	"created_at": "o.created_at",
	"updated_at": "o.updated_at",
}

type OrganizationService struct {
	repo repositories.OrganizationRepository
}

func NewOrganizationService(repo repositories.OrganizationRepository) *OrganizationService {
	return &OrganizationService{repo: repo}
}

func (s *OrganizationService) GetOrgs(search, limit, offset, orderBy, sort string) ([]models.OrganizationWithTotalCount, error) {
	orderBySanitized, sortSanitized, err := internal.SanitizeSort(orderBy, sort, allowedSortColumnsOrgs)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetOrgs(search, limit, offset, orderBySanitized, strings.ToUpper(sortSanitized))
}

func (s *OrganizationService) GetOrg(id string) (models.Organization, error) {
	return s.repo.GetOrg(id)
}

func (s *OrganizationService) DeleteOrg(id string) (models.Organization, error) {
	return s.repo.DeleteOrg(id)
}
