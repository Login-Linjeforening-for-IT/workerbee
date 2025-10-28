// repositories/stats_repositories.go
package repositories

import (
	"encoding/json"
	"log"
	"os"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type Statsrepositories interface {
	GetTotalStats() ([]models.TotalStats, error)
	GetCategoriesStats() ([]models.CategoriesStats, error)
	GetNewAdditionsStats() (models.GroupedNewAdditionsStats, error)
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

func (r *statsrepositories) GetNewAdditionsStats() (models.GroupedNewAdditionsStats, error) {
	var result struct {
		GroupedData []byte `db:"grouped_data"`
	}
	sqlBytes, err := os.ReadFile("./db/stats/get_last_ten_rows_from_tables_grouped.sql")
	if err != nil {
		log.Println("unable to read SQL file:", err)
		return models.GroupedNewAdditionsStats{}, err
	}

	query := string(sqlBytes)
	if err := r.db.Get(&result, query); err != nil {
		log.Println("unable to query DB:", err)
		return models.GroupedNewAdditionsStats{}, err
	}

	var groupedData models.GroupedNewAdditionsStats
	if err := json.Unmarshal(result.GroupedData, &groupedData); err != nil {
		log.Println("unable to unmarshal JSON:", err)
		return models.GroupedNewAdditionsStats{}, err
	}

	return groupedData, nil
}
