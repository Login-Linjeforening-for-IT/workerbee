package models

import "time"

type Event struct {
	ID                 int        `db:"id" json:"id"`
	Visible            bool       `db:"visible" json:"visible"`
	NameNo             string     `db:"name_no" json:"name_no"`
	NameEn             string     `db:"name_en" json:"name_en"`
	DescriptionNo      string     `db:"description_no" json:"description_no"`
	DescriptionEn      string     `db:"description_en" json:"description_en"`
	InformationalNo    string     `db:"informational_no" json:"informational_no"`
	InformationalEn    string     `db:"informational_en" json:"informational_en"`
	TimeType           string     `db:"time_type" json:"time_type"`
	TimeStart          time.Time  `db:"time_start" json:"time_start"`
	TimeEnd            time.Time  `db:"time_end" json:"time_end"`
	TimePublish        *time.Time `db:"time_publish" json:"time_publish,omitempty"`
	TimeSignupRelease  *time.Time `db:"time_signup_release" json:"time_signup_release,omitempty"`
	TimeSignupDeadline *time.Time `db:"time_signup_deadline" json:"time_signup_deadline,omitempty"`
	Canceled           bool       `db:"canceled" json:"canceled"`
	Digital            bool       `db:"digital" json:"digital"`
	Highlight          bool       `db:"highlight" json:"highlight"`
	ImageSmall         *string    `db:"image_small" json:"image_small,omitempty"`
	ImageBanner        string     `db:"image_banner" json:"image_banner"`
	LinkFacebook       *string    `db:"link_facebook" json:"link_facebook,omitempty"`
	LinkDiscord        *string    `db:"link_discord" json:"link_discord,omitempty"`
	LinkSignup         *string    `db:"link_signup" json:"link_signup,omitempty"`
	LinkStream         *string    `db:"link_stream" json:"link_stream,omitempty"`
	Capacity           *int       `db:"capacity" json:"capacity,omitempty"`
	Full               bool       `db:"full" json:"full"`
	Category           int        `db:"category_id" json:"category_id"`
	CategoryNameNo     string     `db:"category_name_no" json:"category_name_no"`
	CategoryNameEn     string     `db:"category_name_en" json:"category_name_en"`
	Location           *int       `db:"location_id" json:"location_id,omitempty"`
	LocationNameNo     string     `db:"location_name_no" json:"location_name_no"`
	LocationNameEn     string     `db:"location_name_en" json:"location_name_en"`
	IsDeleted          bool       `db:"is_deleted" json:"is_deleted"`
	Parent             *int       `db:"parent_id" json:"parent_id,omitempty"`
	Rule               *int       `db:"rule_id" json:"rule_id,omitempty"`
	Audience           *int       `db:"audience_id" json:"audience_id,omitempty"`
	AudienceNameEn     string     `db:"audience_name_en" json:"audience_name_en"`
	AudienceNameNo     string     `db:"audience_name_no" json:"audience_name_no"`
	Organization       *int       `db:"organization_id" json:"organization_id,omitempty"`
	OrganizerNameNo    string     `db:"organizer_name_no" json:"organizer_name_no"`
	OrganizerNameEn    string     `db:"organizer_name_en" json:"organizer_name_en"`
	UpdatedAt          time.Time  `db:"updated_at" json:"updated_at"`
	CreatedAt          time.Time  `db:"created_at" json:"created_at"`
}

type EventWithTotalCount struct {
	Event
	TotalCount int `db:"total_count"`
}

type EventResponse struct {
	Events      []Event `json:"events"`
	TotalCount int     `json:"total_count"`
}
