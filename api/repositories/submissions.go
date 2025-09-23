package repositories

import (
	"encoding/json"
	"os"
	"time"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type Submissionrepositories interface {
	GetSubmission(formID, submissionID string) (*models.Submission, error)
}

type submissionrepositories struct {
	db *sqlx.DB
}

func NewSubmissionrepositories(db *sqlx.DB) Submissionrepositories {
	return &submissionrepositories{db: db}
}

func (r *submissionrepositories) GetSubmission(formID, submissionID string) (*models.Submission, error) {
	type submissionRaw struct {
		ID           int             `db:"id"`
		SubmittedAt  time.Time       `db:"submitted_at"`
		UpdatedAt    time.Time       `db:"updated_at"`
		UserRaw      json.RawMessage `db:"user"`
		QuestionsRaw json.RawMessage `db:"questions"`
	}

	s := submissionRaw{}
	sqlBytes, err := os.ReadFile("./db/forms/submissions/get_submission.sql")
	if err != nil {
		return nil, err
	}

	query := string(sqlBytes)
	err = r.db.Get(&s, query, submissionID, formID)
	if err != nil {
		return nil, err
	}

	sub := &models.Submission{
		ID:          s.ID,
		SubmittedAt: s.SubmittedAt,
		UpdatedAt:   s.UpdatedAt,
	}

	if len(s.UserRaw) > 0 {
		user := models.User{}
		_ = json.Unmarshal(s.UserRaw, &user)
		sub.User = &user
	}

	if len(s.QuestionsRaw) > 0 {
		_ = json.Unmarshal(s.QuestionsRaw, &sub.Questions)
	}

	return sub, nil
}
