package models

import "time"

type Form struct {
	ID          int        `db:"id"`
	UserID      int        `db:"user_id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	Capacity    *int64     `db:"capacity"`
	OpenAt      time.Time  `db:"open_at"`
	CloseAt     time.Time  `db:"close_at"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}

type FormWithTotalCount struct {
	Form
	TotalCount int `db:"total_count"`
}

type FormsResponse struct {
    Forms      []Form `json:"forms"`
    TotalCount int    `json:"total_count"`
}