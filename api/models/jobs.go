package models

import (
	"time"

	"github.com/lib/pq"
)

type Job struct {
	ID                  int            `db:"id"`
	Visible             bool           `db:"visible"`
	Highlight           bool           `db:"highlight"`
	TitleNo             string         `db:"title_no"`
	TitleEn             string         `db:"title_en"`
	Cities              pq.StringArray `db:"cities"`
	Skills              pq.StringArray `db:"skills"`
	PositionTitleNo     string         `db:"position_title_no"`
	PositionTitleEn     string         `db:"position_title_en"`
	DescriptionShortNo  string         `db:"description_short_no"`
	DescriptionShortEn  string         `db:"description_short_en"`
	DescriptionLongNo   string         `db:"description_long_no"`
	DescriptionLongEn   string         `db:"description_long_en"`
	JobType             string         `db:"job_type"`
	TimePublish         time.Time      `db:"time_publish"`
	TimeExpire          time.Time      `db:"time_expire"`
	ApplicationDeadline time.Time      `db:"application_deadline"`
	BannerImage         *string        `db:"banner_image"`
	OrganizationID      int            `db:"organization_id"`
	ApplicationURL      *string        `db:"application_url"`
	CreatedAt           time.Time      `db:"created_at"`
	UpdatedAt           time.Time      `db:"updated_at"`
}

type Cities struct {
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
