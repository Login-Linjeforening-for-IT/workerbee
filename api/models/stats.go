package models

import (
	"workerbee/internal"
)

type YearlyActivity struct {
	InsertDate    internal.Date `db:"insert_date" json:"insert_date"`
	InsertedCount int           `db:"inserted_count" json:"inserted_count"`
}
type CategoriesStats struct {
	ID         int    `db:"id" json:"id"`
	NameEN     string `db:"name_en" json:"name_en"`
	EventCount int    `db:"event_count" json:"event_count"`
	Color      string `db:"color" json:"color"`
}

type NewAddition struct {
	ID        int                `db:"id" json:"id"`
	NameEn    string             `db:"name_en" json:"name_en"`
	CreatedAt internal.LocalTime `db:"created_at" json:"created_at"`
	Source    string             `db:"source" json:"source"`
}
