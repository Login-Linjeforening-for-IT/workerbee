package repositories

import (
	"os"
	"workerbee/db"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type AlertRepository interface {
	GetAlertServices() ([]string, error)
	GetAllPathsInAlertService(service string) ([]models.AlertLanguages, error)
}

type alertRepository struct {
	db *sqlx.DB
}

func NewAlertRepository(db *sqlx.DB) AlertRepository {
	return &alertRepository{db: db}
}

func (r *alertRepository) GetAlertServices() ([]string, error) {
	response, err := db.FetchAllForeignAttributes[string](
		r.db,
		"./db/alerts/get_all_alerts.sql",
	)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *alertRepository) GetAllPathsInAlertService(service string) ([]models.AlertLanguages, error) {
	sqlBytes, err := os.ReadFile("./db/alerts/get_all_paths_in_service.sql")
	if err != nil {
		return nil, err
	}

	var result []models.AlertLanguages
	err = r.db.Select(&result, string(sqlBytes), service)
	if err != nil {
		return nil, err
	}
	return result, nil
}
