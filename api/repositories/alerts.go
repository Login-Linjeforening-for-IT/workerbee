package repositories

import (
	"github.com/jmoiron/sqlx"
)

type AlertRepository interface {
}

type alertRepository struct {
	db *sqlx.DB
}

func NewAlertRepository(db *sqlx.DB) AlertRepository {
	return &alertRepository{db: db}
}
