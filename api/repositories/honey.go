package repositories

import (
	"os"
	"workerbee/db"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type HoneyRepository interface {
	GetTextServices() ([]string, error)
	GetAllPathsInService(service string) ([]models.PathLanguages, error)
	GetAllContentInPath(service, path string) ([]models.HoneyContent, error)
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

func (r *honeyRepository) GetAllPathsInService(service string) ([]models.PathLanguages, error) {
	sqlBytes, err := os.ReadFile("./db/honey/get_all_paths_in_service.sql")
	if err != nil {
		return nil, err
	}

	var result []models.PathLanguages
	err = r.db.Select(&result, string(sqlBytes), service)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *honeyRepository) GetAllContentInPath(service, path string) ([]models.HoneyContent, error) {
	sqlBytes, err := os.ReadFile("./db/honey/get_all_content_in_path.sql")
	if err != nil {
		return nil, err
	}

	var result []models.HoneyContent
	err = r.db.Select(&result, string(sqlBytes), service, path)
	if err != nil {
		return nil, err
	}
	return result, nil
}
