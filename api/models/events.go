package models

import (
	"time"
)

type Event struct {
	ID                 int           `db:"id" json:"id,omitempty"`
	Visible            bool          `db:"visible" json:"visible,omitempty" validate:"omitempty"`
	NameNo             string        `db:"name_no" json:"name_no" validate:"required"`
	NameEn             string        `db:"name_en" json:"name_en" validate:"required"`
	DescriptionNo      string        `db:"description_no" json:"description_no" validate:"required"`
	DescriptionEn      string        `db:"description_en" json:"description_en" validate:"required"`
	InformationalNo    string        `db:"informational_no" json:"informational_no" validate:"required"`
	InformationalEn    string        `db:"informational_en" json:"informational_en" validate:"required"`
	TimeType           string        `db:"time_type" json:"time_type,omitempty" validate:"omitempty"`
	TimeStart          time.Time     `db:"time_start" json:"time_start" validate:"required"`
	TimeEnd            time.Time     `db:"time_end" json:"time_end" validate:"required"`
	TimePublish        time.Time     `db:"time_publish" json:"time_publish" validate:"required"`
	TimeSignupRelease  *time.Time    `db:"time_signup_release" json:"time_signup_release,omitempty" validate:"omitempty"`
	TimeSignupDeadline *time.Time    `db:"time_signup_deadline" json:"time_signup_deadline,omitempty" validate:"omitempty"`
	Canceled           bool          `db:"canceled" json:"canceled,omitempty" validate:"omitempty"`
	Digital            bool          `db:"digital" json:"digital,omitempty" validate:"omitempty"`
	Highlight          bool          `db:"highlight" json:"highlight,omitempty" validate:"omitempty"`
	ImageSmall         *string       `db:"image_small" json:"image_small,omitempty" validate:"omitempty"`
	ImageBanner        string        `db:"image_banner" json:"image_banner" validate:"required"`
	LinkFacebook       *string       `db:"link_facebook" json:"link_facebook,omitempty" validate:"omitempty"`
	LinkDiscord        *string       `db:"link_discord" json:"link_discord,omitempty" validate:"omitempty"`
	LinkSignup         *string       `db:"link_signup" json:"link_signup,omitempty" validate:"omitempty"`
	LinkStream         *string       `db:"link_stream" json:"link_stream,omitempty" validate:"omitempty"`
	Capacity           *int          `db:"capacity" json:"capacity,omitempty" validate:"omitempty"`
	IsFull             bool          `db:"is_full" json:"is_full,omitempty" validate:"omitempty"`
	Category           *Category     `db:"category" json:"category,omitempty" validate:"omitempty"`
	Location           *Location     `db:"location" json:"location,omitempty" validate:"omitempty"`
	ParentID           *int          `db:"parent_id" json:"parent_id,omitempty" validate:"omitempty"`
	Rule               *Rule         `db:"rules" json:"rule,omitempty" validate:"omitempty"`
	Audience           *int          `db:"audience_id" json:"audience_id,omitempty" validate:"omitempty"`
	AudienceNameEn     *string       `db:"audience_name_en" json:"audience_name_en,omitempty" validate:"omitempty"`
	AudienceNameNo     *string       `db:"audience_name_no" json:"audience_name_no,omitempty" validate:"omitempty"`
	Organization       *Organization `db:"organization" json:"organization,omitempty" validate:"omitempty"`
	UpdatedAt          time.Time     `db:"updated_at" json:"updated_at,omitempty" validate:"omitempty"`
	CreatedAt          time.Time     `db:"created_at" json:"created_at,omitempty" validate:"omitempty"`
}

type EventWithTotalCount struct {
	Event
	TotalCount int `db:"total_count"`
}

type EventCategory struct {
	ID     int    `db:"id" json:"id"`
	NameNo string `db:"name_no" json:"name_no"`
	NameEn string `db:"name_en" json:"name_en"`
}
