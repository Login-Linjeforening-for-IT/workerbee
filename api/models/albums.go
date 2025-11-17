package models

import (
	"encoding/json"
	"net/http"
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
	EventID *int `db:"event_id" json:"event_id"`
}

type AlbumWithImages struct {
	Album
	Event      *EventAlbum `db:"event" json:"event"`
	Images     []string    `json:"images"`
	ImageCount int         `db:"image_count" json:"image_count"`
}

type EventAlbum struct {
	ID        *int                `db:"id" json:"id,omitempty"`
	NameEN    *string             `db:"name_en" json:"name_en,omitempty"`
	NameNo    *string             `db:"name_no" json:"name_no,omitempty"`
	TimeStart *internal.LocalTime `db:"time_start" json:"time_start,omitempty"`
	TimeEnd   *internal.LocalTime `db:"time_end" json:"time_end,omitempty"`
}

type AlbumsWithTotalCount struct {
	AlbumWithImages
	TotalCount int `db:"total_count" json:"-"`
}

type UploadImages struct {
	Filename string `json:"filename" validate:"required"`
	Type     string `json:"type" validate:"required"`
}

type UploadPictureResponse struct {
	URL     string      `json:"url"`
	Headers http.Header `json:"headers"`
	Key     string      `json:"key"`
}

func (a AlbumWithImages) MarshalJSON() ([]byte, error) {
	type Alias AlbumWithImages

	aux := &struct {
		*Alias
		Event *EventAlbum `json:"event"`
	}{
		Alias: (*Alias)(&a),
	}

	if a.Event != nil && a.Event.ID == nil {
		aux.Event = nil
	} else {
		aux.Event = a.Event
	}

	return json.Marshal(aux)
}
