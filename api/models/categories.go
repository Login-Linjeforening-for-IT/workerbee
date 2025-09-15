package models

import "time"

type Category struct {
	ID            uint      `db:"id"`
	Color         string    `db:"color"`
	NameNo        string    `db:"name_no"`
	NameEn        string    `db:"name_en"`
	DescriptionNo string    `db:"description_no"`
	DescriptionEn string    `db:"description_en"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type CategoryWithTotalCount struct {
	Category
	TotalCount int `db:"total_count"`
}

type CategoriesResponse struct {
	Categories  []Category `json:"categories"`
	TotalCount  int        `json:"total_count"`
}