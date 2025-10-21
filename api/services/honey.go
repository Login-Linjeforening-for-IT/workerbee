package services

import (
	"encoding/json"
	"workerbee/models"
	"workerbee/repositories"
)

type HoneyService struct {
	repo repositories.HoneyRepository
}

func NewHoneyService(repo repositories.HoneyRepository) *HoneyService {
	return &HoneyService{repo: repo}
}

func (s *HoneyService) GetTextServices() ([]string, error) {
	return s.repo.GetTextServices()
}

func (s *HoneyService) GetAllPathsInService(service string) (map[string][]string, error) {
	rows, err := s.repo.GetAllPathsInService(service)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]string)
	for _, row := range rows {
		result[row.Page] = row.Languages
	}
	return result, nil
}

func (s *HoneyService) GetAllContentInPath(service, path string) (map[string]map[string]string, error) {
	if path[0] != '/' {
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

func (s *HoneyService) GetOneLanguage(service, path, language string) (models.LanguageContentResponse, error) {
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
