package models

import (
	"time"
)

type Location struct {
	ID               int        `db:"id"`
	NameNO           string     `db:"name_no"`
	NameEN           string     `db:"name_en"`
	Type             string     `db:"type"`
	MazemapCampusID  *int       `db:"mazemap_campus_id"`
	MazemapPoiID     *int       `db:"mazemap_poi_id"`
	AddressStreet    *string    `db:"address_street"`
	AddressPostcode  *int       `db:"address_postcode"`
	CityID           *int       `db:"city_id"`
	CoordinateLat    *float64   `db:"coordinate_lat"`
	CoordinateLong   *float64   `db:"coordinate_long"`
	URL              *string    `db:"url"`
	CreatedAt        time.Time  `db:"created_at"`
	UpdatedAt        time.Time  `db:"updated_at"`
}

type LocationWithTotalCount struct {
	Location
	TotalCount int `db:"total_count"`
}

type LocationsResponse struct {
	Locations  []Location `json:"locations"`
	TotalCount int        `json:"total_count"`
}