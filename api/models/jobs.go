package models

import (
	"workerbee/internal"

	"github.com/lib/pq"
)

type BaseJob struct {
	ID                 *int               `db:"id" json:"id"`
	Visible            bool               `db:"visible" json:"visible"`
	Highlight          bool               `db:"highlight" json:"highlight"`
	TitleNo            string             `db:"title_no" json:"title_no"`
	TitleEn            string             `db:"title_en" json:"title_en"`
	Cities             pq.StringArray     `db:"cities" json:"cities"`
	Skills             pq.StringArray     `db:"skills" json:"skills"`
	PositionTitleNo    string             `db:"position_title_no" json:"position_title_no"`
	PositionTitleEn    string             `db:"position_title_en" json:"position_title_en"`
	DescriptionShortNo string             `db:"description_short_no" json:"description_short_no"`
	DescriptionShortEn string             `db:"description_short_en" json:"description_short_en"`
	DescriptionLongNo  string             `db:"description_long_no" json:"description_long_no"`
	DescriptionLongEn  string             `db:"description_long_en" json:"description_long_en"`
	TimePublish        internal.LocalTime `db:"time_publish" json:"time_publish" validate:"required,beforeField=TimeExpire"`
	TimeExpire         internal.LocalTime `db:"time_expire" json:"time_expire" validate:"required,afterField=TimePublish"`
	BannerImage        *string            `db:"banner_image" json:"banner_image,omitempty"`
	ApplicationURL     *string            `db:"application_url" json:"application_url,omitempty"`
	CreatedAt          internal.LocalTime `db:"created_at" json:"created_at"`
	UpdatedAt          internal.LocalTime `db:"updated_at" json:"updated_at"`
}

type Job struct {
	BaseJob
	JobType      JobType      `db:"job_type" json:"job_type"`
	Organization Organization `db:"organization" json:"organization"`
}

type NewJob struct {
	BaseJob
	JobTypeID      int `db:"job_type_id" json:"job_type_id" validate:"required"`
	OrganizationID int `db:"organization_id" json:"organization_id" validate:"required"`
}

type Cities struct {
	ID   *int    `db:"id" json:"id,omitempty"`
	Name *string `db:"name" json:"name,omitempty"`
}

type JobType struct {
	ID        int                `db:"id" json:"id,omitempty"`
	NameNo    string             `db:"name_no" json:"name_no"`
	NameEn    string             `db:"name_en" json:"name_en"`
	UpdatedAt internal.LocalTime `db:"updated_at" json:"updated_at"`
	CreatedAt internal.LocalTime `db:"created_at" json:"created_at"`
}

type JobTypeWithTotalCount struct {
	JobType
	TotalCount int `db:"total_count" json:"-"`
}

type JobSkills struct {
	ID   *int    `db:"id" json:"id,omitempty"`
	Name *string `db:"name" json:"name,omitempty"`
}

type CitiesWithTotalCount struct {
	Cities
	TotalCount int `db:"total_count" json:"-"`
}

type CitiesResponse struct {
	Cities     []CitiesWithTotalCount `json:"cities"`
	TotalCount int                    `json:"total_count"`
}

type JobWithTotalCount struct {
	Job
	TotalCount int `db:"total_count" json:"-"`
}
