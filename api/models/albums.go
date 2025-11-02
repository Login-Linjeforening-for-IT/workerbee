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
	CreatedAt     internal.LocalTime `db:"created_at" json:"created_at"`
	UpdatedAt     internal.LocalTime `db:"updated_at" json:"updated_at"`
}

type CreateAlbum struct {
	Album
	EventID *int `db:"event_id" json:"event_id,omitempty"`
}

type AlbumWithImages struct {
	Album
	Event  *EventAlbum `db:"event" json:"event"`
	Images []string   `json:"images"`
}

type EventAlbum struct {
	ID        *int                `db:"id" json:"id,omitempty"`
	NameEN    *string             `db:"name_en" json:"name_en,omitempty"`
	NameNo    *string             `db:"name_no" json:"name_no,omitempty"`
	TimeStart *internal.LocalTime `db:"time_start" json:"time_start,omitempty"`
	TimeEnd   *internal.LocalTime `db:"time_end" json:"time_end,omitempty"`
}
