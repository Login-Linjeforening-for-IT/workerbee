// services/stats_service.go
package services

import (
	"strconv"
	"workerbee/models"
	"workerbee/repositories"
)

type StatsService struct {
	repo repositories.Statsrepositories
}

func NewStatsService(repo repositories.Statsrepositories) *StatsService {
	return &StatsService{repo: repo}
}

func (s *StatsService) GetMostActiveCategories() ([]models.CategoriesStats, error) {
	return s.repo.GetMostActiveCategories()
}

func (s *StatsService) GetYearlyStats() ([]models.YearlyActivity, error) {
	return s.repo.GetYearlyStats()
}

func (s *StatsService) GetNewAdditionsStats(limit_str string) ([]models.NewAddition, error) {
	limit, err := strconv.Atoi(limit_str)
	if err != nil || limit < 1 {
		limit = 10
	} else if limit > 25 {
		limit = 25
	}

	return s.repo.GetNewAdditionsStats(limit)
}
