package token

import (
	"testing"
	"time"

	"gitlab.login.no/tekkom/web/beehive/admin-api/utils/require"
)

func Test__NewPasetoMaker(t *testing.T) {
	t.Run("Invalid key size", func(t *testing.T) {
		_, err := NewPasetoMaker("invalid-key", AccessToken, 0)
		// require.ErrorAs(t, err, &InvalidKeySizeError{})
		require.NotNil(t, err) // TODO: fix
	})

	t.Run("Valid key", func(t *testing.T) {
		_, err := NewPasetoMaker("01234567890123456789012345678901", AccessToken, 0)
		require.Nil(t, err)
	})

	t.Run("Invalid token type", func(t *testing.T) {
		_, err := NewPasetoMaker("01234567890123456789012345678901", "invalid-token-type", 0)
		require.ErrorIs(t, err, ErrInvalidTokenType)
	})
}

func Test__CreateToken(t *testing.T) {
	maker, err := NewPasetoMaker("01234567890123456789012345678901", AccessToken, 0)
	require.Nil(t, err)

	token, _, err := maker.CreateToken(CreateTokenParams{
		UID:   "user1",
		Roles: []string{"role1", "role2"},
	})
	require.Nil(t, err)
	require.NotEqual(t, token, "")
}

func Test__VerifyToken(t *testing.T) {
	maker, err := NewPasetoMaker(
		"01234567890123456789012345678901",
		AccessToken,
		15*time.Minute,
	)
	require.Nil(t, err)

	token, _, err := maker.CreateToken(CreateTokenParams{
		UID:   "user1",
		Roles: []string{"role1", "role2"},
	})
	require.Nil(t, err)

	payload, err := maker.VerifyToken(token)
	require.Nil(t, err)
	require.Equal(t, payload.UID, "user1")
	require.Equal(t, payload.Roles, []string{"role1", "role2"})

	// expired token
	maker, err = NewPasetoMaker(
		"01234567890123456789012345678901",
		AccessToken,
		-1*time.Second,
	)
	require.Nil(t, err)

	token, _, err = maker.CreateToken(CreateTokenParams{
		UID:   "user1",
		Roles: []string{"role1", "role2"},
	})
	require.Nil(t, err)

	payload, err = maker.VerifyToken(token)
	require.ErrorIs(t, err, ErrExpiredToken)
	require.Nil(t, payload)
}
