package models

import (
	"time"
	"workerbee/internal"
)

type User struct {
	ID        int       `db:"id" json:"id"`
	FullName  string    `db:"full_name" json:"full_name"`
	Email     string    `db:"email" json:"email"`
	CreatedAt internal.LocalTime `db:"created_at" json:"created_at"`
	UpdatedAt internal.LocalTime `db:"updated_at" json:"updated_at"`
}

type Form struct {
	ID          int       `db:"id" json:"id"`
	UserID      int       `db:"user_id" json:"user_id"`
	Title       string    `db:"title" json:"title"`
	Description *string   `db:"description" json:"description"`
	Capacity    *int      `db:"capacity" json:"capacity"`
	OpenAt      internal.LocalTime `db:"open_at" json:"open_at"`
	CloseAt     internal.LocalTime `db:"close_at" json:"close_at"`
	CreatedAt   internal.LocalTime `db:"created_at" json:"created_at"`
	UpdatedAt   internal.LocalTime `db:"updated_at" json:"updated_at"`
}

type FormWithTotalCount struct {
	Form
	TotalCount int `db:"total_count"`
}

type FormsResponse struct {
    Forms      []FormWithTotalCount `json:"forms"`
    TotalCount int                  `json:"total_count"`
}

type FormWithQuestion struct {
	ID          int                    `db:"id" json:"id"`
	Title       string                 `db:"title" json:"title"`
	Description *string                `db:"description" json:"description"`
	Capacity    *int                   `db:"capacity" json:"capacity"`
	OpenAt      time.Time              `db:"open_at" json:"open_at"`
	CloseAt     time.Time              `db:"close_at" json:"close_at"`
	CreatedAt   time.Time              `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time              `db:"updated_at" json:"updated_at"`
	User        *User                  `db:"user" json:"user"`
	Questions   []QuestionWithOptions  `db:"questions" json:"questions"`
}