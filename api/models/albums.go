package models

import (
	"workerbee/internal"
)

type Album struct {
	ID            int                `db:"id"`
	NameEn        string             `db:"name_en"`
	NameNo        string             `db:"name_no"`
	DescriptionEn string             `db:"description_en"`
	DescriptionNo string             `db:"description_no"`
	CreatedAt     internal.LocalTime `db:"created_at"`
	UpdatedAt     internal.LocalTime `db:"updated_at"`
}
