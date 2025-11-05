package services

import (
	"encoding/json"
	"strconv"
	"strings"
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repositories"
)

var validLanguages = []string{
	"en",
	"no",
}

type HoneyService struct {
	repo repositories.HoneyRepository
}

func NewHoneyService(repo repositories.HoneyRepository) *HoneyService {
	return &HoneyService{repo: repo}
}

func (s *HoneyService) CreateHoney(
	content models.CreateHoney,
) (models.CreateHoney, error) {
	return s.repo.CreateTextInService(content)
}

func (s *HoneyService) GetTextServices() ([]string, error) {
	return s.repo.GetTextServices()
}

func (s *HoneyService) GetAllPathsInService(service string) ([]models.PathLanguagesWithCount, error) {
	return s.repo.GetAllPathsInService(service)
}

func (s *HoneyService) GetAllContentInPath(service, path string) (map[string]map[string]string, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	rows, err := s.repo.GetAllContentInPath(service, path)
	if err != nil {
		return nil, err
	}

	result := make(map[string]map[string]string)
	for _, row := range rows {
		var content map[string]string
		err := json.Unmarshal([]byte(row.Text), &content)
		if err != nil {
			return nil, err
		}
		result[row.Language] = content
	}
	return result, nil
}

func (s *HoneyService) UpdateContentInPath(id_str string, content models.CreateHoney) (models.CreateHoney, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.CreateHoney{}, internal.ErrInvalid
	}

	content.ID = id

	return s.repo.UpdateContentInPath(content)
}

func (s *HoneyService) GetOneLanguage(service, path, language string) (models.LanguageContentResponse, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	row, err := s.repo.GetOneLanguage(service, path, language)
	if err != nil {
		return models.LanguageContentResponse{}, err
	}

	formattedText := make(map[string]map[string]string)
	var content map[string]string

	err = json.Unmarshal([]byte(row.Text), &content)
	if err != nil {
		return models.LanguageContentResponse{}, err
	}
	formattedText[row.Language] = content

	return models.LanguageContentResponse{
		Service:  row.Service,
		Page:     row.Page,
		Language: row.Language,
		Text:     formattedText,
	}, nil
}

func (s *HoneyService) DeleteHoney(id string) (int, error) {
	return s.repo.DeleteHoney(id)
}