package models

import (
	"workerbee/internal"
)

type Question struct {
	ID                  int                `db:"id" json:"id"`
	FormID              int                `db:"form_id" json:"form_id"`
	QuestionTitle       string             `db:"question_title" json:"question_title"`
	QuestionDescription string             `db:"question_description" json:"question_description"`
	QuestionType        string             `db:"question_type" json:"question_type"`
	Required            bool               `db:"required" json:"required"`
	Position            int                `db:"position" json:"position"`
	Max                 *int               `db:"max" json:"max"`
	CreatedAt           internal.LocalTime `db:"created_at" json:"created_at"`
	UpdatedAt           internal.LocalTime `db:"updated_at" json:"updated_at"`
}

type QuestionOption struct {
	ID         int    `db:"id" json:"id"`
	OptionText string `db:"option_text" json:"option_text"`
	Position   int    `db:"position" json:"position"`
}

type QuestionWithOptions struct {
	ID                  int              `db:"id" json:"id"`
	QuestionTitle       string           `db:"question_title" json:"question_title"`
	QuestionDescription string           `db:"question_description" json:"question_description"`
	QuestionType        string           `db:"question_type" json:"question_type"`
	Required            bool             `db:"required" json:"required"`
	Position            int              `db:"position" json:"position"`
	Max                 *int             `db:"max" json:"max"`
	Options             []QuestionOption `db:"options" json:"options"`
}
