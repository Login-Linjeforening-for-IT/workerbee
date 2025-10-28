package models

import (
	"workerbee/internal"
)

type YearlyActivity struct {
	InsertDate internal.LocalTime `db:"insert_date" json:"insert_date"`
	InsertedCount int                `db:"inserted_count" json:"inserted_count"`
}
type CategoriesStats struct {
	ID         int    `db:"id" json:"id"`
	NameEN     string `db:"name_en" json:"name_en"`
	EventCount int    `db:"event_count" json:"event_count"`
}

type newAdditionsStats struct {
	ID        int                `json:"id"`
	CreatedAt internal.LocalTime `json:"created_at"`
	NameEN    string             `json:"name_en"`
}

type GroupedNewAdditionsStats struct {
	Categories    []newAdditionsStats `json:"categories"`
	Events        []newAdditionsStats `json:"events"`
	Locations     []newAdditionsStats `json:"locations"`
	Jobs          []newAdditionsStats `json:"jobs"`
	Audiences     []newAdditionsStats `json:"audiences"`
	Rules         []newAdditionsStats `json:"rules"`
	Organizations []newAdditionsStats `json:"organizations"`
}
