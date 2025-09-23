package models

import (
	"time"
)

type Organization struct {
	ID            int       `db:"id" json:"id"`
	ShortName     string    `db:"shortname" json:"short_name"`
	NameNo        string    `db:"name_no" json:"name_no"`
	NameEn        string    `db:"name_en" json:"name_en"`
	DescriptionNo string    `db:"description_no" json:"description_no"`
	DescriptionEn string    `db:"description_en" json:"description_en"`
	Type          int       `db:"type" json:"type"`
	LinkHomepage  string    `db:"link_homepage" json:"link_homepage"`
	LinkLinkedin  string    `db:"link_linkedin" json:"link_linkedin"`
	LinkFacebook  string    `db:"link_facebook" json:"link_facebook"`
	LinkInstagram string    `db:"link_instagram" json:"link_instagram"`
	Logo          string    `db:"logo" json:"logo"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

type OrganizationWithTotalCount struct {
	Organization
	TotalCount int `db:"total_count"`
}

type OrganizationsResponse struct {
	Organizations []OrganizationWithTotalCount `json:"organizations"`
	TotalCount    int                          `json:"total_count"`
}