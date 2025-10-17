package services

import (
	"strconv"
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repositories"
)

var allowedSortColumnsAudiences = map[string]string{
	"id":      "a.id",
	"name_no": "a.name_no",
	"name_en": "a.name_en",
}

type AudienceService struct {
	repo repositories.Audiencerepository
}

func NewAudienceService(repo repositories.Audiencerepository) *AudienceService {
	return &AudienceService{repo: repo}
}

func (s *AudienceService) CreateAudience(audience models.Audience) (models.Audience, error) {
	return s.repo.CreateAudience(audience)
}

func (s *AudienceService) GetAudience(id string) (models.Audience, error) {
	return s.repo.GetAudience(id)
}

func (s *AudienceService) GetAudiences(search, limit_str, offset_str, orderBy, sort string) ([]models.AudienceWithTotalCount, error) {
	orderBySanitized, sortSanitized, err := internal.SanitizeSort(orderBy, sort, allowedSortColumnsAudiences)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetAudiences(limit, offset, search, orderBySanitized, sortSanitized)
}

func (s *AudienceService) UpdateAudience(id_str string, audience models.Audience) (models.Audience, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.Audience{}, internal.ErrInvalid
	}
	audience.ID = &id
	return s.repo.UpdateAudience(audience)
}

func (s *AudienceService) DeleteAudience(id string) (int, error) {
	return s.repo.DeleteAudience(id)
}
