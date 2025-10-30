package repositories

import (
	"workerbee/db"

	"github.com/jmoiron/sqlx"
)

type AlertRepository interface {
	GetAlertServices() ([]string, error)
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
