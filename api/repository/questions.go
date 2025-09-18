package repository

import (
	"os"

	"github.com/jmoiron/sqlx"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
)

type QuestionRepository interface {
	PostQuestions(formID string, questions []models.Question) ([]models.Question, error)
	PatchQuestions(formID string, questions []models.Question) ([]models.Question, error)
	DeleteQuestion(id string) (models.Question, error)
	PostQuestionOption(questionID string, options models.QuestionOption) (models.QuestionOption, error)
	PatchQuestionOption(options models.QuestionOption) (models.QuestionOption, error)
	DeleteQuestionOption(id string) (models.QuestionOption, error)
}

type questionRepository struct {
	db *sqlx.DB
}

func NewQuestionRepository(db *sqlx.DB) QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) PostQuestions(formID string, questions []models.Question) ([]models.Question, error) {
	inserted := []models.Question{}
	
	sqlBytes, err := os.ReadFile("./db/forms/questions/post_question.sql")
	if err != nil {
		return nil, err
	}
	
	query := string(sqlBytes)
	for _, q := range questions {
		nq := models.Question{}
		err := r.db.Get(&nq, query, formID, q.QuestionTitle, q.QuestionDescription, q.QuestionType, q.Required, q.Position, q.Max)
		if err != nil {
			return nil, err
		}
		inserted = append(inserted, nq)
	}
	
	return inserted, nil
}

func (r *questionRepository) PatchQuestions(formID string, questions []models.Question) ([]models.Question, error) {
	updated := []models.Question{}
	
	sqlBytes, err := os.ReadFile("./db/forms/questions/patch_question.sql")
	if err != nil {
		return nil, err
	}
	
	query := string(sqlBytes)
	for _, q := range questions {
		uq := models.Question{}
		err := r.db.Get(&uq, query, q.ID, q.QuestionTitle, q.QuestionDescription, q.QuestionType, q.Required, q.Position, q.Max)
		if err != nil {
			return nil, err
		}
		updated = append(updated, uq)
	}
	
	return updated, nil
}

func (r *questionRepository) DeleteQuestion(id string) (models.Question, error) {
	question := models.Question{}
	
	sqlBytes, err := os.ReadFile("./db/forms/questions/delete_question.sql")
	if err != nil {
		return question, err
	}
	
	query := string(sqlBytes)
	err = r.db.Get(&question, query, id)
	if err != nil {
		return question, err
	}
	
	return question, nil
}

func (r *questionRepository) PostQuestionOption(questionID string, options models.QuestionOption) (models.QuestionOption, error) {
	option := models.QuestionOption{}
	
	sqlBytes, err := os.ReadFile("./db/forms/questions/post_question_option.sql")
	if err != nil {
		return option, err
	}
	
	query := string(sqlBytes)
	err = r.db.Get(&option, query, questionID, options.OptionText, options.Position)
	if err != nil {
		return option, err
	}
	
	return option, nil
}

func (r *questionRepository) PatchQuestionOption(options models.QuestionOption) (models.QuestionOption, error) {
	option := models.QuestionOption{}
	
	sqlBytes, err := os.ReadFile("./db/forms/questions/patch_question_option.sql")
	if err != nil {
		return option, err
	}
	
	query := string(sqlBytes)
	err = r.db.Get(&option, query, options.ID, options.OptionText, options.Position)
	if err != nil {
		return option, err
	}
	
	return option, nil
}

func (r *questionRepository) DeleteQuestionOption(id string) (models.QuestionOption, error) {
	option := models.QuestionOption{}
	
	sqlBytes, err := os.ReadFile("./db/forms/questions/delete_question_option.sql")
	if err != nil {
		return option, err
	}
	
	query := string(sqlBytes)
	err = r.db.Get(&option, query, id)
	if err != nil {
		return option, err
	}
	
	return option, nil
}