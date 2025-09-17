package services

import (
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/repository"
)

type QuestionService struct {
	repo repository.QuestionRepository
}

func NewQuestionService(repo repository.QuestionRepository) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) GetQuestions(formID string) ([]models.QuestionWithOptions, error) {
	return s.repo.GetQuestions(formID)
}

func (s *QuestionService) PostQuestions(formID string, questions []models.Question) ([]models.Question, error) {
	return s.repo.PostQuestions(formID, questions)
}

func (s *QuestionService) PatchQuestions(formID string, questions []models.Question) ([]models.Question, error) {
	return s.repo.PatchQuestions(formID, questions)
}

func (s *QuestionService) DeleteQuestion(id string) (models.Question, error) {
	return s.repo.DeleteQuestion(id)
}
