package models

import (
	"time"

	"github.com/lib/pq"
)

type Job struct {
	ID                  int            `db:"id" json:"id"`
	Visible             bool           `db:"visible" json:"visible"`
	Highlight           bool           `db:"highlight" json:"highlight"`
	TitleNo             string         `db:"title_no" json:"title_no"`
	TitleEn             string         `db:"title_en" json:"title_en"`
	Cities              pq.StringArray `db:"cities" json:"cities"`
	Skills              pq.StringArray `db:"skills" json:"skills"`
	PositionTitleNo     string         `db:"position_title_no" json:"position_title_no"`
	PositionTitleEn     string         `db:"position_title_en" json:"position_title_en"`
	DescriptionShortNo  string         `db:"description_short_no" json:"description_short_no"`
	DescriptionShortEn  string         `db:"description_short_en" json:"description_short_en"`
	DescriptionLongNo   string         `db:"description_long_no" json:"description_long_no"`
	DescriptionLongEn   string         `db:"description_long_en" json:"description_long_en"`
	JobType             string         `db:"job_type" json:"job_type"`
	TimePublish         time.Time      `db:"time_publish" json:"time_publish"`
	TimeExpire          time.Time      `db:"time_expire" json:"time_expire"`
	ApplicationDeadline time.Time      `db:"application_deadline" json:"application_deadline"`
	BannerImage         *string        `db:"banner_image" json:"banner_image,omitempty"`
	Organization        *Organization  `db:"organization" json:"organization,omitempty"`
	ApplicationURL      *string        `db:"application_url" json:"application_url,omitempty"`
	CreatedAt           time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt           time.Time      `db:"updated_at" json:"updated_at"`
}


type Cities struct {
	ID   int    `db:"id" json:"id,omitempty"`
	Name string `db:"name"`
}

type JobType struct {
	Type string `db:"job_type"`
}

type JobSkills struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type CitiesWithTotalCount struct {
	Cities
	TotalCount int `db:"total_count"`
}

type CitiesResponse struct {
	Jobs       []CitiesWithTotalCount `json:"cities"`
	TotalCount int                    `json:"total_count"`
}

type JobWithTotalCount struct {
	Job
	TotalCount int `db:"total_count"`
}

type JobsResponse struct {
	Jobs       []JobWithTotalCount `json:"jobs"`
	TotalCount int                 `json:"total_count"`
}
