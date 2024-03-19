package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken     = errors.New("token is invalid")
	ErrExpiredToken     = errors.New("token has expired")
	ErrInvalidTokenType = errors.New("invalid token type")
)

type tokenType string

const (
	AccessToken  tokenType = "access"
	RefreshToken tokenType = "refresh"
)

func (t tokenType) Valid() bool {
	switch t {
	case AccessToken, RefreshToken:
		return true
	}
	return false
}

type Payload struct {
	ID        string    `json:"id"`
	UID       string    `json:"uid"`
	Roles     []string  `json:"roles"`
	Type      tokenType `json:"type"`
	IssuedAt  time.Time `json:"iat"`
	ExpiresAt time.Time `json:"exp"`
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}

	return nil
}

func NewPayload(
	params CreateTokenParams,
	tokenType tokenType,
	duration time.Duration,
) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	issuedAt := time.Now()

	return &Payload{
		ID:        id.String(),
		UID:       params.UID,
		Roles:     params.Roles,
		Type:      tokenType,
		IssuedAt:  issuedAt,
		ExpiresAt: issuedAt.Add(duration),
	}, nil
}
