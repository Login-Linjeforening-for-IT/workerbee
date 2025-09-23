package models

import (
	"time"
)

type Location struct {
	ID              int       `db:"id" json:"id"`
	NameNO          string    `db:"name_no" json:"name_no"`
	NameEN          string    `db:"name_en" json:"name_en"`
	Type            string    `db:"type" json:"type"`
	MazemapCampusID *int      `db:"mazemap_campus_id" json:"mazemap_campus_id,omitempty"`
	MazemapPoiID    *int      `db:"mazemap_poi_id" json:"mazemap_poi_id,omitempty"`
	AddressStreet   *string   `db:"address_street" json:"address_street,omitempty"`
	AddressPostcode *int      `db:"address_postcode" json:"address_postcode,omitempty"`
	CityID          *int      `db:"city_id" json:"city_id,omitempty"`
	CoordinateLat   *float64  `db:"coordinate_lat" json:"coordinate_lat,omitempty"`
	CoordinateLong  *float64  `db:"coordinate_long" json:"coordinate_long,omitempty"`
	URL             *string   `db:"url" json:"url,omitempty"`
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
