package models

import (
	"time"
)

type Organization struct {
	ID             int        `db:"id"`
	ShortName      string     `db:"shortname"`
	NameNo         string     `db:"name_no"`
	NameEn         string     `db:"name_en"`
	DescriptionNo  string     `db:"description_no"`
	DescriptionEn  string     `db:"description_en"`
	Type           int        `db:"type"`
	LinkHomepage   string     `db:"link_homepage"`
	LinkLinkedin   string     `db:"link_linkedin"`
	LinkFacebook   string     `db:"link_facebook"`
	LinkInstagram  string     `db:"link_instagram"`
	Logo           string     `db:"logo"`
	CreatedAt      time.Time  `db:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at"`
}

type OrganizationWithTotalCount struct {
	Organization
	TotalCount int `db:"total_count"`
}

type OrganizationsResponse struct {
	Organizations []OrganizationWithTotalCount `json:"organizations"`
	TotalCount    int                          `json:"total_count"`
}