package sessionstore

import "time"

type Session struct {
	ID  string `json:"id"`
	UID string `json:"uid"`

	RefreshToken string `json:"refresh_token"`
	UserAgent    string `json:"user_agent"`
	ClientIP     string `json:"client_ip"`

	ExpiresAt time.Time `json:"exp"`
	BlockedAt time.Time `json:"blocked_at"`
	CreatedAt time.Time `json:"created_at"`
}
