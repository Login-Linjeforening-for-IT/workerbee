package models

import "workerbee/internal"

type CalendarEvent struct {
	ID            int                `db:"id" json:"id"`
	NameNo        string             `db:"name_no" json:"name_no"`
	NameEn        string             `db:"name_en" json:"name_en"`
	DescriptionEn string             `db:"description_en" json:"description_en"`
	DescriptionNo string             `db:"description_no" json:"description_no"`
	TimeStart     internal.LocalTime `db:"time_start" json:"time_start"`
	TimeEnd       internal.LocalTime `db:"time_end" json:"time_end"`
}
