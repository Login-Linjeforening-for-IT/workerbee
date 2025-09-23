package models

import "time"

type Category struct {
	ID            uint      `json:"id" db:"id"`
	Color         string    `json:"color" db:"color"`
	NameNo        string    `json:"name_no" db:"name_no"`
	NameEn        string    `json:"name_en" db:"name_en"`
	DescriptionNo string    `json:"description_no" db:"description_no"`
	DescriptionEn string    `json:"description_en" db:"description_en"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type CategoryWithTotalCount struct {
	Category
	TotalCount int `db:"total_count"`
}

type CategoriesResponse struct {
	Categories  []Category `json:"categories"`
	TotalCount  int        `json:"total_count"`
}