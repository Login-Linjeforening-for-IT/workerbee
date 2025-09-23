package services

import (
	"strings"
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repositories"
)

var allowedSortColumnsRules = map[string]string{
	"id":         "r.id",
	"name_no":    "r.name_no",
	"name_en":    "r.name_en",
	"created_at": "r.created_at",
	"updated_at": "r.updated_at",
}

type RuleService struct {
	repo repositories.Rulerepositories
}

func NewRuleService(repo repositories.Rulerepositories) *RuleService {
	return &RuleService{repo: repo}
}

func (s *RuleService) GetRules(search, limit, offset, orderBy, sort string) ([]models.RuleWithTotalCount, error) {
	orderBySanitized, sortSanitized, err := internal.SanitizeSort(orderBy, sort, allowedSortColumnsRules)
	if err != nil {
		return nil, internal.ErrInvalid
	}
	return s.repo.GetRules(search, limit, offset, orderBySanitized, strings.ToUpper(sortSanitized))
}
