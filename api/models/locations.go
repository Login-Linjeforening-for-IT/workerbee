package models

import (
	"time"
)

type Location struct {
	ID              int       `db:"id" json:"id,omitempty"`
	NameNO          string    `db:"name_no" json:"name_no" validate:"required"`
	NameEN          string    `db:"name_en" json:"name_en" validate:"required"`
	Type            string    `db:"type" json:"type"`
	MazemapCampusID *int      `db:"mazemap_campus_id" json:"mazemap_campus_id,omitempty"`
	MazemapPoiID    *int      `db:"mazemap_poi_id" json:"mazemap_poi_id,omitempty"`
	AddressStreet   *string   `db:"address_street" json:"address_street,omitempty"`
	AddressPostcode *int      `db:"address_postcode" json:"address_postcode,omitempty"`
	City            *Cities   `db:"cities" json:"cities,omitempty"`
	CoordinateLat   *float64  `db:"coordinate_lat" json:"coordinate_lat,omitempty"`
	CoordinateLon   *float64  `db:"coordinate_lon" json:"coordinate_lon,omitempty"`
	URL             *string   `db:"url" json:"url,omitempty" validate:"omitempty,url"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}

type LocationWithTotalCount struct {
	Location
	TotalCount int `db:"total_count"`
}

type LocationsResponse struct {
	Locations  []Location `json:"locations"`
	TotalCount int        `json:"total_count"`
}
