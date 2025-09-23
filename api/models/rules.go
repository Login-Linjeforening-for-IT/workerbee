package models

import "time"

type Rule struct {
	ID            int       `db:"id" json:"id"`
	NameNo        string    `db:"name_no" json:"name_no"`
	NameEn        string    `db:"name_en" json:"name_en"`
	DescriptionNo string    `db:"description_no" json:"description_no"`
	DescriptionEn string    `db:"description_en" json:"description_en"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

type RuleWithTotalCount struct {
	Rule
	TotalCount int `db:"total_count"`
}

type RulesResponse struct {
	Rules      []Rule `json:"rules"`
	TotalCount int    `json:"total_count"`
}
