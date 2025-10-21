package repositories

import (
	"workerbee/db"

	"github.com/jmoiron/sqlx"
)

type HoneyRepository interface {
	GetTextServices() ([]string, error)
}

type honeyRepository struct {
	db *sqlx.DB
}

func NewHoneyRepository(db *sqlx.DB) HoneyRepository {
	return &honeyRepository{db: db}
}

func (r *honeyRepository) GetTextServices() ([]string, error) {
	response, err := db.FetchAllForeignAttributes[string](
		r.db,
		"./db/honey/get_all_services.sql",
	)
	if err != nil {
		return nil, err
	}
	return response, nil
}
