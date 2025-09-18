package repository

import (
	"encoding/json"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
)

type SubmissionRepository interface {
	GetSubmission(formID, submissionID string) (*models.Submission, error)
}

type submissionRepository struct {
	db *sqlx.DB
}

func NewSubmissionRepository(db *sqlx.DB) SubmissionRepository {
	return &submissionRepository{db: db}
}

func (r *submissionRepository) GetSubmission(formID, submissionID string) (*models.Submission, error) {
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
