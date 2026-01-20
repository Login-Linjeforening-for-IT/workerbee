package models

import (
	"workerbee/internal"
)

type Organization struct {
	ID            *int                `db:"id" json:"id,omitempty"`
	NameNo        *string             `db:"name_no" json:"name_no" validate:"required"`
	NameEn        *string             `db:"name_en" json:"name_en" validate:"required"`
	DescriptionNo *string             `db:"description_no" json:"description_no" validate:"required"`
	DescriptionEn *string             `db:"description_en" json:"description_en" validate:"required"`
	LinkHomepage  *string             `db:"link_homepage" json:"link_homepage" validate:"required"`
	LinkLinkedin  *string             `db:"link_linkedin" json:"link_linkedin"`
	LinkFacebook  *string             `db:"link_facebook" json:"link_facebook"`
	LinkInstagram *string             `db:"link_instagram" json:"link_instagram"`
	Logo          *string             `db:"logo" json:"logo"`
	CreatedAt     *internal.LocalTime `db:"created_at" json:"created_at"`
	UpdatedAt     *internal.LocalTime `db:"updated_at" json:"updated_at"`
}

type OrganizationWithTotalCount struct {
	Organization
	TotalCount int `db:"total_count" json:"-"`
}

type OrganizationsResponse struct {
	Organizations []OrganizationWithTotalCount `json:"organizations"`
	TotalCount    int                          `json:"-"`
}

type OrganizationNames struct {
	ID     int    `db:"id" json:"id"`
	NameNo string `db:"name_no" json:"name_no"`
	NameEn string `db:"name_en" json:"name_en"`
}
