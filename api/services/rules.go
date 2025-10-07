package services

import (
	"strconv"
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

func (s *RuleService) GetRule(id string) (models.Rule, error) {
	return s.repo.GetRule(id)
}

func (s *RuleService) DeleteRule(id string) (models.Rule, error) {
	return s.repo.DeleteRule(id)
}

func (s *RuleService) CreateRule(rule models.Rule) (models.Rule, error) {
	return s.repo.CreateRule(rule)
}

func (s *RuleService) UpdateRule(id_str string, rule models.Rule) (models.Rule, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.Rule{}, internal.ErrInvalid
	}

	rule.ID = id

	return s.repo.UpdateRule(rule)
}
