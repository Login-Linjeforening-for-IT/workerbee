package services

import (
	"workerbee/models"
	"workerbee/repository"
)

type RuleService struct {
	repo repository.RuleRepository
}

func NewRuleService(repo repository.RuleRepository) *RuleService {
	return &RuleService{repo: repo}
}

func (s *RuleService) GetRules(search, limit, offset, orderBy, sort string) ([]models.RuleWithTotalCount, error) {
	return s.repo.GetRules(search, limit, offset, orderBy, sort)
}
