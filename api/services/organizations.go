package services

import (
	"strconv"
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

func (s *OrganizationService) GetOrgs(search, limit_str, offset_str, orderBy, sort string) ([]models.OrganizationWithTotalCount, error) {
	orderBySanitized, sortSanitized, err := internal.SanitizeSort(orderBy, sort, allowedSortColumnsOrgs)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetOrgs(limit, offset, search, orderBySanitized, strings.ToUpper(sortSanitized))
}

func (s *OrganizationService) GetOrg(id string) (models.Organization, error) {
	return s.repo.GetOrg(id)
}

func (s *OrganizationService) DeleteOrg(id string) (models.Organization, error) {
	return s.repo.DeleteOrg(id)
}

func (s *OrganizationService) UpdateOrg(id_str string, org models.Organization) (models.Organization, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.Organization{}, internal.ErrInvalid
	}

	org.ID = &id

	return s.repo.UpdateOrg(org)
}

func (s *OrganizationService) CreateOrg(org models.Organization) (models.Organization, error) {
	return s.repo.CreateOrg(org)
}
