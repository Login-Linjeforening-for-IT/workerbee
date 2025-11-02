// repositories/stats_repositories.go
package repositories

import (
	"log"
	"os"
	"workerbee/db"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type Statsrepositories interface {
	GetYearlyStats() ([]models.YearlyActivity, error)
	GetCategoriesStats() ([]models.CategoriesStats, error)
	GetNewAdditionsStats() ([]models.NewAddition, error)
	GetMostActiveCategories() ([]models.CategoriesStats, error)
}

type statsrepositories struct {
	db *sqlx.DB
}

func NewStatsrepositories(db *sqlx.DB) Statsrepositories {
	return &statsrepositories{db: db}
}

func (r *statsrepositories) GetMostActiveCategories() ([]models.CategoriesStats, error) {
	return db.FetchAllForeignAttributes[models.CategoriesStats](
		r.db,
		"./db/stats/get_most_active_categories.sql",
	)
}

func (r *statsrepositories) GetYearlyStats() ([]models.YearlyActivity, error) {
	stats, err := db.FetchAllForeignAttributes[models.YearlyActivity](
		r.db,
		"./db/stats/get_yearly_stats.sql",
	)
	if err != nil {
		return nil, err
	}
	return stats, nil
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

func (r *statsrepositories) GetNewAdditionsStats() ([]models.NewAddition, error) {
	return db.FetchAllForeignAttributes[models.NewAddition](
		r.db,
		"./db/stats/get_new_additions_stats.sql",
	)
}
