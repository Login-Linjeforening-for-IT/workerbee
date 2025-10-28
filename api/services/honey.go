package services

import (
	"encoding/json"
	"slices"
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

func (s *HoneyService) UpdateContentInPath(service, path string, content map[string]map[string]string) (map[string]any, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	if len(content) != 2 {
		return nil, internal.ErrInvalid
	}

	for lang := range content {
		if !slices.Contains(validLanguages, lang) {
			return nil, internal.ErrInvalid
		}
	}

	err := validateHoneyContent(content)
	if err != nil {
		return nil, err
	}

	return s.repo.UpdateContentInPath(service, path, content)
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

func validateHoneyContent(content map[string]map[string]string) error {
	enFields := make(map[string]bool)
	noFields := make(map[string]bool)
		
	for field := range content["en"] {
		enFields[field] = true
	}
	for field := range content["no"] {
		noFields[field] = true
	}
	
	if len(enFields) != len(noFields) {
		return internal.ErrInvalid
	}
	
	for field := range enFields {
		if !noFields[field] {
			return internal.ErrInvalid
		}
	}
	return nil
}