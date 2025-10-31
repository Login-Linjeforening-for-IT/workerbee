package repositories

import (
	"workerbee/db"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type AlertRepository interface {
	GetAllAlerts(limit, offset int, search, orderBy, sort string) ([]models.Alert, error)
	GetAlertByServiceAndPage(service, page string) (models.Alert, error)
}

type alertRepository struct {
	db *sqlx.DB
}

func NewAlertRepository(db *sqlx.DB) AlertRepository {
	return &alertRepository{db: db}
}

func (r *alertRepository) GetAllAlerts(limit, offset int, search, orderBy, sort string) ([]models.Alert, error) {
	return db.FetchAllElements[models.Alert](
		r.db,
		"./db/alerts/get_all_alerts.sql",
		orderBy, sort,
		limit,
		offset,
		search,
	)
}

func (r *alertRepository) GetAlertByServiceAndPage(service, page string) (models.Alert, error) {
	alert, err := db.ExecuteOneRow[models.Alert](
		r.db,
		"./db/alerts/get_alert_by_service_and_page.sql",
		service,
		page,
	)
	return alert, err
}
