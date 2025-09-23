package services

import (
	"workerbee/models"
	"workerbee/repositories"
)

type QuestionService struct {
	repo repositories.Questionrepositories
}

func NewQuestionService(repo repositories.Questionrepositories) *QuestionService {
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
