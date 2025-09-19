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

func (s *QuestionService) PostQuestions(formID string, questions []models.Question) ([]models.Question, error) {
	return s.repo.PostQuestions(formID, questions)
}

func (s *QuestionService) PutQuestions(formID string, questions []models.Question) ([]models.Question, error) {
	return s.repo.PutQuestions(formID, questions)
}

func (s *QuestionService) DeleteQuestion(id string) (models.Question, error) {
	return s.repo.DeleteQuestion(id)
}

func (s *QuestionService) PostQuestionOption(questionID string, options models.QuestionOption) (models.QuestionOption, error) {
	return s.repo.PostQuestionOption(questionID, options)
}

func (s *QuestionService) PutQuestionOption(options models.QuestionOption) (models.QuestionOption, error) {
	return s.repo.PutQuestionOption(options)
}

func (s *QuestionService) DeleteQuestionOption(id string) (models.QuestionOption, error) {
	return s.repo.DeleteQuestionOption(id)
}
