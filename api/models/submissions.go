package models

import (
	"time"
)

type Submission struct {
	ID          int                  `db:"id" json:"id"`
	SubmittedAt time.Time            `db:"submitted_at" json:"submitted_at"`
	UpdatedAt   time.Time            `db:"updated_at" json:"updated_at"`
	User        *User                `db:"user" json:"user"`
	Questions   []QuestionWithAnswer `db:"questions" json:"questions"`
}

type Answer struct {
	ID              int     `db:"id" json:"id"`
	AnswerText      *string `db:"answer_text" json:"answer_text"`
	SelectedOptions []int   `db:"selected_options" json:"selected_options"`
}

type QuestionWithAnswer struct {
	ID                  int              `db:"id" json:"id"`
	QuestionTitle       string           `db:"question_title" json:"question_title"`
	QuestionDescription string           `db:"question_description" json:"question_description"`
	QuestionType        string           `db:"question_type" json:"question_type"`
	Required            bool             `db:"required" json:"required"`
	Position            int              `db:"position" json:"position"`
	Max                 *int             `db:"max" json:"max"`
	Options             []QuestionOption `db:"options" json:"options"`
	Answer              []Answer         `db:"answer" json:"answer"`
}
