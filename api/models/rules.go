package models

import "time"

type Rule struct {
	ID            int        `db:"id"`
	NameNo        string     `db:"name_no"`
	NameEn        string     `db:"name_en"`
	DescriptionNo string     `db:"description_no"`
	DescriptionEn string     `db:"description_en"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at"`
}

type RuleWithTotalCount struct {
	Rule
	TotalCount int `db:"total_count"`
}

type RulesResponse struct {
	Rules      []Rule `json:"rules"`
	TotalCount int    `json:"total_count"`
}