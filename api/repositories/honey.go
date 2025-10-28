package repositories

import (
	"encoding/json"
	"errors"
	"os"
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type HoneyRepository interface {
	CreateTextInService(service, path, language string, content map[string]map[string]string) (map[string]any, error)
	GetTextServices() ([]string, error)
	GetAllPathsInService(service string) ([]models.PathLanguages, error)
	GetAllContentInPath(service, path string) ([]models.HoneyContent, error)
	GetOneLanguage(service, path, language string) (models.LanguageContent, error)
	UpdateContentInPath(service, path string, content map[string]map[string]string) (map[string]any, error)
}

type honeyRepository struct {
	db *sqlx.DB
}

func NewHoneyRepository(db *sqlx.DB) HoneyRepository {
	return &honeyRepository{db: db}
}

func (r *honeyRepository) CreateTextInService(
	service, path, language string,
	content map[string]map[string]string,
) (map[string]any, error) {
	var pqErr *pq.Error
	sqlBytes, err := os.ReadFile("./db/honey/add_content_in_service.sql")
	if err != nil {
		return nil, err
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var languages []string

	for language, fields := range content {
		contentJSON, err := json.Marshal(fields)
		if err != nil {
			return nil, err
		}

		_, err = tx.Exec(string(sqlBytes), string(contentJSON), service, path, language)
		if err != nil {
			if errors.As(err, &pqErr) && pqErr.Code == "23505" {
				return nil, internal.ErrConflict
			}
			return nil, err
		}

		languages = append(languages, language)
	}

	resp := make(map[string]any)
	resp["status"] = "success"
	resp["service"] = service
	resp["path"] = path
	resp["updated"] = languages

	return resp, nil
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

func (r *honeyRepository) GetOneLanguage(service, path, language string) (models.LanguageContent, error) {
	result, err := db.ExecuteOneRow[models.LanguageContent](
		r.db,
		"./db/honey/get_info_for_one_language.sql",
		service, path, language,
	)
	if err != nil {
		return models.LanguageContent{}, err
	}
	return result, nil
}

func (r *honeyRepository) UpdateContentInPath(service, path string, content map[string]map[string]string) (map[string]any, error) {
	sqlBytes, err := os.ReadFile("./db/honey/update_content_in_path.sql")
	if err != nil {
		return nil, err
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var languages []string

	for language, fields := range content {
		contentJSON, err := json.Marshal(fields)
		if err != nil {
			return nil, err
		}

		_, err = tx.Exec(string(sqlBytes), string(contentJSON), service, path, language)
		if err != nil {
			return nil, err
		}

		languages = append(languages, language)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	resp := make(map[string]any)
	resp["status"] = "success"
	resp["service"] = service
	resp["path"] = path
	resp["updated"] = languages

	return resp, nil
}
