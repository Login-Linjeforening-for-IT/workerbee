package models

import "workerbee/internal"

type Alert struct {
	ID            int                `db:"id" json:"id,omitempty"`
	Service       string             `db:"service" json:"service" validate:"required"`
	Page          string             `db:"page" json:"page" validate:"required"`
	TitleEn       string             `db:"title_en" json:"title_en" validate:"required"`
	TitleNo       string             `db:"title_no" json:"title_no" validate:"required"`
	Description   string             `db:"description_en" json:"description_en" validate:"required"`
	DescriptionNo string             `db:"description_no" json:"description_no" validate:"required"`
	UpdatedAt     internal.LocalTime `db:"updated_at" json:"updated_at,omitempty"`
	CreatedAt     internal.LocalTime `db:"created_at" json:"created_at,omitempty"`
}

type AlertWithTotalCount struct {
	Alert
	TotalCount int `db:"total_count" json:"total_count"`
}