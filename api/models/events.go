package models

import "time"

type Event struct {
	ID                 int        `db:"id"`
	Visible            bool       `db:"visible"`
	NameNo             string     `db:"name_no"`
	NameEn             string     `db:"name_en"`
	DescriptionNo      string     `db:"description_no"`
	DescriptionEn      string     `db:"description_en"`
	InformationalNo    string     `db:"informational_no"`
	InformationalEn    string     `db:"informational_en"`
	TimeType           string     `db:"time_type"` // Use string or custom type for enum
	TimeStart          time.Time  `db:"time_start"`
	TimeEnd            time.Time  `db:"time_end"`
	TimePublish        *time.Time `db:"time_publish"`
	TimeSignupRelease  *time.Time `db:"time_signup_release"`
	TimeSignupDeadline *time.Time `db:"time_signup_deadline"`
	Canceled           bool       `db:"canceled"`
	Digital            bool       `db:"digital"`
	Highlight          bool       `db:"highlight"`
	ImageSmall         *string    `db:"image_small"`
	ImageBanner        string     `db:"image_banner"`
	LinkFacebook       *string    `db:"link_facebook"`
	LinkDiscord        *string    `db:"link_discord"`
	LinkSignup         *string    `db:"link_signup"`
	LinkStream         *string    `db:"link_stream"`
	Capacity           *int       `db:"capacity"`
	Full               bool       `db:"full"`
	Category           int        `db:"category"`
	CategoryNameNo     string     `db:"category_name_no"`
	CategoryNameEn     string     `db:"category_name_en"`
	Location           *int       `db:"location"`
	Location_name_no   string     `db:"location_name_no"`
	Location_name_en   string     `db:"location_name_en"`
	Is_deleted         bool       `db:"is_deleted"`
	Parent             *int       `db:"parent"`
	Rule               *int       `db:"rule"`
	Audience           *int       `db:"audience"`
	Audience_name_en   string     `db:"audience_name_en"`
	Audience_name_no   string     `db:"audience_name_no"`
	Organization       *int       `db:"organization"`
	Organizer_name_no  string     `db:"organizer_name_no"`
	Organizer_name_en  string     `db:"organizer_name_en"`
	UpdatedAt          time.Time  `db:"updated_at"`
	CreatedAt          time.Time  `db:"created_at"`
	DeletedAt          *time.Time `db:"deleted_at"`
}
