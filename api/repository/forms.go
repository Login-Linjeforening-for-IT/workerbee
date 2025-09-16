package repository

import (
	"os"

	"github.com/jmoiron/sqlx"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
)

type FormRepository interface {
	GetForms(search, limit, offset string) ([]models.FormWithTotalCount, error)
	GetForm(id string) ([]models.Form, error)
}

type formRepository struct {
	db *sqlx.DB
}

func NewFormRepository(db *sqlx.DB) FormRepository {
	return &formRepository{db: db}
}

func (r *formRepository) GetForms(search, limit, offset string) ([]models.FormWithTotalCount, error) {
	forms := []models.FormWithTotalCount{}

	sqlBytes, err := os.ReadFile("./db/forms/get_forms.sql")
	if err != nil {
		return nil, err
	}

	query := string(sqlBytes)
	err = r.db.Select(&forms, query, search, limit, offset)
	if err != nil {
		return nil, err
	}

	return forms, nil
}

func (r *formRepository) GetForm(id string) ([]models.Form, error) {
	forms := []models.Form{}

	sqlBytes, err := os.ReadFile("./db/forms/get_form.sql")
	if err != nil {
		return nil, err
	}

	query := string(sqlBytes)
	err = r.db.Select(&forms, query, id)
	if err != nil {
		return nil, err
	}
	return forms, nil
}
