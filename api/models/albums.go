package models

import (
	"workerbee/internal"
)

type Album struct {
	ID            int                `db:"id" json:"id"`
	NameEn        string             `db:"name_en" json:"name_en"`
	NameNo        string             `db:"name_no" json:"name_no"`
	DescriptionEn string             `db:"description_en" json:"description_en"`
	DescriptionNo string             `db:"description_no" json:"description_no"`
	Year          int                `db:"year" json:"year"`
	EventID       *int               `db:"event_id" json:"event_id"`
	CreatedAt     internal.LocalTime `db:"created_at" json:"created_at"`
	UpdatedAt     internal.LocalTime `db:"updated_at" json:"updated_at"`
}
