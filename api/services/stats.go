// services/stats_service.go
package services

import (
	"workerbee/models"
	"workerbee/repositories"
)

type StatsService struct {
	repo repositories.Statsrepositories
}

func NewStatsService(repo repositories.Statsrepositories) *StatsService {
	return &StatsService{repo: repo}
}

func (s *StatsService) GetMostActiveCategory() (models.CategoriesStats, error) {
	return s.repo.GetMostActiveCategory()
}

func (s *StatsService) GetYearlyStats() ([]models.YearlyActivity, error) {
	return s.repo.GetYearlyStats()
}

func (s *StatsService) GetCategoriesStats() ([]models.CategoriesStats, error) {
	return s.repo.GetCategoriesStats()
}

func (s *StatsService) GetNewAdditionsStats() (models.GroupedNewAdditionsStats, error) {
	return s.repo.GetNewAdditionsStats()
}
