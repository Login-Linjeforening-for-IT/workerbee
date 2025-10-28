// repositories/stats_repositories.go
package repositories

import (
	"log"
	"os"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type Statsrepositories interface {
	GetTotalStats() ([]models.TotalStats, error)
	GetCategoriesStats() ([]models.CategoriesStats, error)
	GetNewAdditionsStats(limit int) ([]models.NewAdditionsStats, error)
	GetMostActiveCategory() (models.CategoriesStats, error)
}

type statsrepositories struct {
	db *sqlx.DB
}

func NewStatsrepositories(db *sqlx.DB) Statsrepositories {
	return &statsrepositories{db: db}
}

func (r *statsrepositories) GetMostActiveCategory() (models.CategoriesStats, error) {
	var categoryStat models.CategoriesStats
	sqlBytes, err := os.ReadFile("./db/stats/get_most_active_category.sql")
	if err != nil {
		return models.CategoriesStats{}, err
	}

	query := string(sqlBytes)
	if err := r.db.Get(&categoryStat, query); err != nil {
		return models.CategoriesStats{}, err
	}

	return categoryStat, nil
}

func (r *statsrepositories) GetTotalStats() ([]models.TotalStats, error) {
	totalStats := []models.TotalStats{}
	sqlBytes, err := os.ReadFile("./db/stats/get_total_stats.sql")
	if err != nil {
		return nil, err
	}

	query := string(sqlBytes)
	if err := r.db.Select(&totalStats, query); err != nil {
		return nil, err
	}

	return totalStats, nil
}

func (r *statsrepositories) GetCategoriesStats() ([]models.CategoriesStats, error) {
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

func (r *statsrepositories) GetNewAdditionsStats(limit int) ([]models.NewAdditionsStats, error) {
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
