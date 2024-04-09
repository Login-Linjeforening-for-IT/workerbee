package sessionstore

import (
	"context"
	"time"
)

type Store interface {
	CreateSession(ctx context.Context, params CreateSessionParams) error
	GetSession(ctx context.Context, id string) (Session, error)
	BlockSession(ctx context.Context, id string) error
	DeleteSession(ctx context.Context, id string) error
}

type CreateSessionParams struct {
	ID           string
	UID          string
	RefreshToken string
	UserAgent    string
	ClientIP     string
	ExpiresAt    time.Time
}
