// services/stats_service.go
package services

import (
	"workerbee/models"
	"workerbee/repository"
)

type StatsService struct {
	repo repository.StatsRepository
}

func NewStatsService(repo repository.StatsRepository) *StatsService {
	return &StatsService{repo: repo}
}

func (s *StatsService) GetTotalStats() ([]models.TotalStats, error) {
	return s.repo.GetTotalStats()
}

func (s *StatsService) GetCategoriesStats() ([]models.CategoriesStats, error) {
	return s.repo.GetCategoriesStats()
}

func (s *StatsService) GetNewAdditionsStats(limit int) ([]models.NewAdditionsStats, error) {
	return s.repo.GetNewAdditionsStats(limit)
}
