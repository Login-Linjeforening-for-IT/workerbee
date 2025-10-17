package models

import "workerbee/internal"

type Category struct {
	ID        *int               `json:"id" db:"id"`
	NameNo    string             `json:"name_no" db:"name_no" validate:"required"`
	NameEn    string             `json:"name_en" db:"name_en" validate:"required"`
	Color     string             `json:"color" db:"color" validate:"required"`
	CreatedAt internal.LocalTime `json:"created_at" db:"created_at"`
	UpdatedAt internal.LocalTime `json:"updated_at" db:"updated_at"`
}

type CategoryWithTotalCount struct {
	Category
	TotalCount int `db:"total_count"`
}
