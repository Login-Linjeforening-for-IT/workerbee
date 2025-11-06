package repositories

import (
	"os"
	"workerbee/db"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type HoneyRepository interface {
	CreateTextInService(content models.CreateHoney) (models.CreateHoney, error)
	GetTextServices() ([]string, error)
	GetAllPathsInService(service string) ([]models.PathLanguagesWithCount, error)
	GetAllContentInPath(service, path string) ([]models.HoneyContent, error)
	GetOneLanguage(service, path, language string) (models.CreateHoney, error)
	GetHoney(id int) (models.CreateHoney, error)
	UpdateContentInPath(content models.CreateHoney) (models.CreateHoney, error)
	DeleteHoney(id string) (int, error)
}

type honeyRepository struct {
	db *sqlx.DB
}

func NewHoneyRepository(db *sqlx.DB) HoneyRepository {
	return &honeyRepository{db: db}
}

func (r *honeyRepository) CreateTextInService(
	content models.CreateHoney,
) (models.CreateHoney, error) {
	return db.AddOneRow(
		r.db,
		"./db/honey/add_service_with_content.sql",
		content,
	)
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

func (r *honeyRepository) GetAllPathsInService(service string) ([]models.PathLanguagesWithCount, error) {
	sqlBytes, err := os.ReadFile("./db/honey/get_all_paths_in_service.sql")
	if err != nil {
		return nil, err
	}

	var result []models.PathLanguagesWithCount
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

func (r *honeyRepository) GetOneLanguage(service, path, language string) (models.CreateHoney, error) {
	result, err := db.ExecuteOneRow[models.CreateHoney](
		r.db,
		"./db/honey/get_info_for_one_language.sql",
		service, path, language,
	)
	if err != nil {
		return models.CreateHoney{}, err
	}
	return result, nil
}

func (r *honeyRepository) GetHoney(id int) (models.CreateHoney, error) {
	result, err := db.ExecuteOneRow[models.CreateHoney](
		r.db,
		"./db/honey/get_one_honey.sql",
		id,
	)
	if err != nil {
		return models.CreateHoney{}, err
	}
	return result, nil
}

func (r *honeyRepository) UpdateContentInPath(content models.CreateHoney) (models.CreateHoney, error) {
	updatedContent, err := db.AddOneRow(
		r.db,
		"./db/honey/update_content_in_path.sql",
		content,
	)
	if err != nil {
		return models.CreateHoney{}, err
	}
	return updatedContent, nil
}

func (r *honeyRepository) DeleteHoney(id string) (int, error) {
	honeyID, err := db.ExecuteOneRow[int](
		r.db,
		"./db/honey/delete_honey.sql",
		id,
	)
	if err != nil {
		return 0, err
	}
	return honeyID, nil
}
