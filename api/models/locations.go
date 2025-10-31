package models

import "workerbee/internal"

type LocationBase struct {
	ID              *int                `db:"id" json:"id,omitempty"`
	NameNO          *string             `db:"name_no" json:"name_no" validate:"required"`
	NameEN          *string             `db:"name_en" json:"name_en" validate:"required"`
	Type            *string             `db:"type" json:"type"`
	MazemapCampusID *int                `db:"mazemap_campus_id" json:"mazemap_campus_id,omitempty"`
	MazemapPoiID    *int                `db:"mazemap_poi_id" json:"mazemap_poi_id,omitempty"`
	AddressStreet   *string             `db:"address_street" json:"address_street,omitempty"`
	AddressPostcode *int                `db:"address_postcode" json:"address_postcode,omitempty"`
	CoordinateLat   *float64            `db:"coordinate_lat" json:"coordinate_lat,omitempty" validate:"omitempty,gte=-90,lte=90"`
	CoordinateLon   *float64            `db:"coordinate_lon" json:"coordinate_lon,omitempty" validate:"omitempty,gte=-180,lte=180"`
	URL             *string             `db:"url" json:"url,omitempty"`
	CreatedAt       *internal.LocalTime `db:"created_at" json:"created_at"`
	UpdatedAt       *internal.LocalTime `db:"updated_at" json:"updated_at"`
}

type Location struct {
	LocationBase
	City *Cities `db:"cities" json:"cities,omitempty"`
}

type NewLocation struct {
	LocationBase
	CityID *int `db:"city_id" json:"city_id"`
}

type LocationWithTotalCount struct {
	Location
	TotalCount int `db:"total_count"`
}

type LocationNames struct {
	ID     int    `db:"id" json:"id"`
	NameEN string `db:"name_en" json:"name_en"`
	Type   string `db:"type" json:"type"`
}