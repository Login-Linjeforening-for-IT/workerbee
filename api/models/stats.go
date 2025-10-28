package models

import (
	"workerbee/internal"
)

type TotalStats struct {
	TotalEvents        int `db:"total_events"`
	TotalJobs          int `db:"total_jobs"`
	TotalOrganizations int `db:"total_organizations"`
	TotalLocations     int `db:"total_locations"`
	TotalRules         int `db:"total_rules"`
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
