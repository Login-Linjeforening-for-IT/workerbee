package repositories

import (
	"workerbee/db"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type AlertRepository interface {
	CreateAlert(alert models.Alert) (models.Alert, error)
	GetAllAlerts(limit, offset int, search, orderBy, sort string) ([]models.AlertWithTotalCount, error)
	GetAlertByServiceAndPage(service, page string) (models.Alert, error)
	GetAlertByID(id int) (models.Alert, error)
	UpdateAlert(alert models.Alert) (models.Alert, error)
	DeleteAlert(id string) (int, error)
}

type alertRepository struct {
	db *sqlx.DB
}

func NewAlertRepository(db *sqlx.DB) AlertRepository {
	return &alertRepository{db: db}
}

func (r *alertRepository) CreateAlert(alert models.Alert) (models.Alert, error) {
	return db.AddOneRow(
		r.db,
		"./db/alerts/post_alert.sql",
		alert,
	)
}

func (r *alertRepository) GetAllAlerts(limit, offset int, search, orderBy, sort string) ([]models.AlertWithTotalCount, error) {
	return db.FetchAllElements[models.AlertWithTotalCount](
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

func (r *alertRepository) GetAlertByID(id int) (models.Alert, error) {
	alert, err := db.ExecuteOneRow[models.Alert](
		r.db,
		"./db/alerts/get_alert_by_id.sql",
		id,
	)
	return alert, err
}

func (r *alertRepository) UpdateAlert(alert models.Alert) (models.Alert, error) {
	return db.AddOneRow(
		r.db,
		"./db/alerts/update_alert.sql",
		alert,
	)
}

func (r *alertRepository) DeleteAlert(id string) (int, error) {
	return db.ExecuteOneRow[int](
		r.db,
		"./db/alerts/delete_alert.sql",
		id,
	)
}
