package api

import (
	"context"
	"fmt"

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

func getAccessTokenFromHeader(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		return "", fmt.Errorf("authorization header not found")
	}

	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return "", fmt.Errorf("invalid authorization format")
	}

	accessToken := authHeader[7:]

	return accessToken, nil
}

// func setRefreshTokenCookie(ctx *gin.Context, refreshToken string, refreshTokenPayload *token.Payload) {
// 	ctx.SetCookie(refreshTokenCookieName, refreshToken,
// 		int(refreshTokenPayload.ExpiresAt.Sub(refreshTokenPayload.IssuedAt).Seconds()), "", "", false, true)
// }

func getRefreshTokenFromHeader(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("X-Refresh-Token")

	if authHeader == "" {
		return "", fmt.Errorf("refresh token header not found")
	}

	return authHeader, nil
}
