package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte

	tokenType tokenType
	duration  time.Duration
}

type InvalidKeySizeError struct {
	Expected int
	Actual   int
}

func (err *InvalidKeySizeError) Error() string {
	return fmt.Sprintf("invalid symmetric key size: expected %d, got %d", err.Expected, err.Actual)
}

func NewPasetoMaker(symmetricKey string, tokenType tokenType, duration time.Duration) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, &InvalidKeySizeError{
			Expected: chacha20poly1305.KeySize,
			Actual:   len(symmetricKey),
		}
	}

	if !tokenType.Valid() {
		return nil, ErrInvalidTokenType
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
		tokenType:    tokenType,
		duration:     duration,
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(params CreateTokenParams) (string, *Payload, error) {
	payload, err := maker.NewPayload(params)
	if err != nil {
		return "", payload, err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)

	return token, payload, err
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	if payload.Type != maker.tokenType {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (maker *PasetoMaker) NewPayload(params CreateTokenParams) (*Payload, error) {
	return NewPayload(params, maker.tokenType, maker.duration)
}
