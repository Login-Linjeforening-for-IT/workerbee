package token

import (
	"testing"
	"time"

	"git.logntnu.no/tekkom/web/beehive/admin-api/utils/require"
)

func Test__NewPayload(t *testing.T) {
	payload, err := NewPayload(CreateTokenParams{
		UID:   "user1",
		Roles: []string{"role1", "role2"},
	}, AccessToken, time.Minute*15)

	require.Nil(t, err)
	require.Equal(t, "user1", payload.UID)
	require.Equal(t, []string{"role1", "role2"}, payload.Roles)
	require.Equal(t, AccessToken, payload.Type)
	require.True(t, time.Now().Before(payload.ExpiresAt))
	require.True(t, time.Now().Add(time.Minute*15).After(payload.ExpiresAt))
}

func Test__Payload_Valid(t *testing.T) {
	payload, err := NewPayload(CreateTokenParams{
		UID:   "user1",
		Roles: []string{"role1", "role2"},
	}, AccessToken, time.Minute*15)

	require.Nil(t, err)
	require.Nil(t, payload.Valid())

	payload, err = NewPayload(CreateTokenParams{
		UID:   "user1",
		Roles: []string{"role1", "role2"},
	}, AccessToken, -time.Minute*15)

	require.Nil(t, err)
	require.ErrorIs(t, payload.Valid(), ErrExpiredToken)
}
