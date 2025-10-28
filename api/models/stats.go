package models

import "time"

type TotalStats struct {
	TotalEvents        int `db:"total_events"`
	TotalJobs          int `db:"total_jobs"`
	TotalOrganizations int `db:"total_organizations"`
	TotalLocations     int `db:"total_locations"`
	TotalRules         int `db:"total_rules"`
}

type CategoriesStats struct {
	ID int    `db:"id" json:"id"`
	NameEN     string `db:"name_en" json:"name_en"`
	EventCount int    `db:"event_count" json:"event_count"`
}

type NewAdditionsStats struct {
	ID        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	Table     string    `db:"table"`
	Name      string    `db:"name"`
}