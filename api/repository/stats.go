// repository/stats_repository.go
package repository

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
)

type StatsRepository interface {
	GetTotalStats() ([]models.TotalStats, error)
	GetCategoriesStats() ([]models.CategoriesStats, error)
	GetNewAdditionsStats(limit int) ([]models.NewAdditionsStats, error)
}

type statsRepository struct {
	db *sqlx.DB
}

func NewStatsRepository(db *sqlx.DB) StatsRepository {
	return &statsRepository{db: db}
}

func (r *statsRepository) GetTotalStats() ([]models.TotalStats, error) {
	totalStats := []models.TotalStats{}
	sqlBytes, err := os.ReadFile("./db/stats/get_total_stats.sql")
	if err != nil {
		log.Println("unable to read SQL file:", err)
		return nil, err
	}

	query := string(sqlBytes)
	if err := r.db.Select(&totalStats, query); err != nil {
		log.Println("unable to query DB:", err)
		return nil, err
	}

	return totalStats, nil
}

func (r *statsRepository) GetCategoriesStats() ([]models.CategoriesStats, error) {
	categoriesStats := []models.CategoriesStats{}
	sqlBytes, err := os.ReadFile("./db/stats/get_categories_stats.sql")
	if err != nil {
		log.Println("unable to read SQL file:", err)
		return nil, err
	}

	query := string(sqlBytes)
	if err := r.db.Select(&categoriesStats, query); err != nil {
		log.Println("unable to query DB:", err)
		return nil, err
	}

	return categoriesStats, nil
}

func (r *statsRepository) GetNewAdditionsStats(limit int) ([]models.NewAdditionsStats, error) {
	newAdditions := []models.NewAdditionsStats{}
	sqlBytes, err := os.ReadFile("./db/stats/get_new_additions_stats.sql")
	if err != nil {
		log.Println("unable to read SQL file:", err)
		return nil, err
	}

	query := string(sqlBytes)
	if err := r.db.Select(&newAdditions, query, limit); err != nil {
		log.Println("unable to query DB:", err)
		return nil, err
	}

	return newAdditions, nil
}
