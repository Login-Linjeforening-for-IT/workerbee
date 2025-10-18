package models

import (
	"workerbee/internal"
)

type Rule struct {
	ID            *int                `db:"id" json:"id"`
	NameNo        *string             `db:"name_no" json:"name_no" validate:"required"`
	NameEn        *string             `db:"name_en" json:"name_en" validate:"required"`
	DescriptionNo *string             `db:"description_no" json:"description_no" validate:"required"`
	DescriptionEn *string             `db:"description_en" json:"description_en" validate:"required"`
	CreatedAt     *internal.LocalTime `db:"created_at" json:"created_at"`
	UpdatedAt     *internal.LocalTime `db:"updated_at" json:"updated_at"`
}

type RuleWithTotalCount struct {
	Rule
	TotalCount int `db:"total_count"`
}

type RulesResponse struct {
	Rules      []Rule `json:"rules"`
	TotalCount int    `json:"total_count"`
}

type RuleNames struct {
	NameEn string `db:"name_en" json:"name_en"`
}
