package api

import (
	"context"

	"gitlab.login.no/tekkom/web/beehive/admin-api/token"

	"github.com/gin-gonic/gin"
)

const (
	accessTokenCookieName  = "access_token"
	refreshTokenCookieName = "refresh_token"
)

// _ = ctx, but was unused
func (server *Server) createAccessToken(_ context.Context, userID string, roles []string) (string, *token.Payload, error) {
	return server.accessTokenMaker.CreateToken(token.CreateTokenParams{
		UID:   userID,
		Roles: roles,
	})
}

// _ = ctx, but was unused
func (server *Server) createRefreshToken(_ context.Context, userID string, roles []string) (string, *token.Payload, error) {
	return server.refreshTokenMaker.CreateToken(token.CreateTokenParams{
		UID:   userID,
		Roles: roles,
	})
}

func setAccessTokenCookie(ctx *gin.Context, accessToken string, accessTokenPayload *token.Payload) {
	ctx.SetCookie(accessTokenCookieName, accessToken,
		int(accessTokenPayload.ExpiresAt.Sub(accessTokenPayload.IssuedAt).Seconds()), "", "", false, true)
}

func getAccessTokenCookie(ctx *gin.Context) (string, error) {
	return ctx.Cookie(accessTokenCookieName)
}

func setRefreshTokenCookie(ctx *gin.Context, refreshToken string, refreshTokenPayload *token.Payload) {
	ctx.SetCookie(refreshTokenCookieName, refreshToken,
		int(refreshTokenPayload.ExpiresAt.Sub(refreshTokenPayload.IssuedAt).Seconds()), "", "", false, true)
}

func getRefreshTokenCookie(ctx *gin.Context) (string, error) {
	return ctx.Cookie(refreshTokenCookieName)
}
