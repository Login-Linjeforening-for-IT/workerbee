package models

import "time"

type User struct {
	ID        int       `db:"id"`
	FullName  string    `db:"full_name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Form struct {
	ID          int       `db:"id"`
	UserID      int       `db:"user_id"`
	Title       string    `db:"title"`
	Description *string   `db:"description"`
	Capacity    *int      `db:"capacity"`
	OpenAt      time.Time `db:"open_at"`
	CloseAt     time.Time `db:"close_at"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Question struct {
	ID                 int       `db:"id"`
	FormID             int       `db:"form_id"`
	QuestionTitle      string    `db:"question_title"`
	QuestionDescription string   `db:"question_description"`
	QuestionType       string    `db:"question_type"`
	Required           bool      `db:"required"`
	Position           int       `db:"position"`
	Max                *int      `db:"max"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

type QuestionOption struct {
	ID           int       `db:"id"`
	QuestionID   int       `db:"question_id"`
	OptionText   string    `db:"option_text"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type QuestionWithOptions struct {
	Question
	Options []QuestionOption `db:"options"`
}

type Submission struct {
	ID          int       `db:"id"`
	FormID      int       `db:"form_id"`
	UserID      int       `db:"user_id"`
	SubmittedAt time.Time `db:"submitted_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Answer struct {
	ID           int       `db:"id"`
	SubmissionID int       `db:"submission_id"`
	QuestionID   int       `db:"question_id"`
	OptionID     *int      `db:"option_id"`
	AnswerText   *string   `db:"answer_text"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type AnswerOption struct {
	AnswerID  int `db:"answer_id"`
	OptionID  int `db:"option_id"`
}

type FormWithTotalCount struct {
	Form
	TotalCount int `db:"total_count"`
}

type FormsResponse struct {
    Forms      []FormWithTotalCount `json:"forms"`
    TotalCount int                  `json:"total_count"`
}