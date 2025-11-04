package repositories

import (
	"os"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type Questionrepositories interface {
	PostQuestions(formID string, questions []models.Question) ([]models.Question, error)
	PutQuestions(formID string, questions []models.Question) ([]models.Question, error)
	DeleteQuestion(id string) (int, error)
	PostQuestionOption(questionID string, options models.QuestionOption) (models.QuestionOption, error)
	PutQuestionOption(options models.QuestionOption) (models.QuestionOption, error)
	DeleteQuestionOption(id string) (int, error)
}

type questionrepositories struct {
	db *sqlx.DB
}

func NewQuestionrepositories(db *sqlx.DB) Questionrepositories {
	return &questionrepositories{db: db}
}

func (r *questionrepositories) PostQuestions(formID string, questions []models.Question) ([]models.Question, error) {
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

func (r *questionrepositories) PutQuestions(formID string, questions []models.Question) ([]models.Question, error) {
	updated := []models.Question{}

	sqlBytes, err := os.ReadFile("./db/forms/questions/put_question.sql")
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

func (r *questionrepositories) DeleteQuestion(id string) (int, error) {
	var questionId int

	sqlBytes, err := os.ReadFile("./db/forms/questions/delete_option.sql")
	if err != nil {
		return 0, err
	}

	query := string(sqlBytes)
	err = r.db.Get(&questionId, query, id)
	if err != nil {
		return 0, err
	}

	return questionId, nil
}

func (r *questionrepositories) PostQuestionOption(questionID string, options models.QuestionOption) (models.QuestionOption, error) {
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

func (r *questionrepositories) PutQuestionOption(options models.QuestionOption) (models.QuestionOption, error) {
	option := models.QuestionOption{}

	sqlBytes, err := os.ReadFile("./db/forms/questions/put_question_option.sql")
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

func (r *questionrepositories) DeleteQuestionOption(id string) (int, error) {
	var optionId int

	sqlBytes, err := os.ReadFile("./db/forms/questions/delete_option.sql")
	if err != nil {
		return 0, err
	}

	query := string(sqlBytes)
	err = r.db.Get(&optionId, query, id)
	if err != nil {
		return 0, err
	}

	return optionId, nil
}
