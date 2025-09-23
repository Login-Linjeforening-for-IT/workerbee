package services

import (
	"workerbee/models"
	"workerbee/repositories"
)

type RuleService struct {
	repo repositories.Rulerepositories
}

func NewRuleService(repo repositories.Rulerepositories) *RuleService {
	return &RuleService{repo: repo}
}

func (s *RuleService) GetRules(search, limit, offset, orderBy, sort string) ([]models.RuleWithTotalCount, error) {
	return s.repo.GetRules(search, limit, offset, orderBy, sort)
}
