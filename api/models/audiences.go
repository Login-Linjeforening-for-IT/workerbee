package models

import "workerbee/internal"

type Audience struct {
	ID        *int                `db:"id" json:"id"`
	NameEn    *string             `db:"name_en" json:"name_en" validate:"required"`
	NameNo    *string             `db:"name_no" json:"name_no" validate:"required"`
	UpdatedAt *internal.LocalTime `db:"updated_at" json:"updated_at"`
	CreatedAt *internal.LocalTime `db:"created_at" json:"created_at"`
}

type AudienceWithTotalCount struct {
	Audience
	TotalCount int `db:"total_count"`
}
