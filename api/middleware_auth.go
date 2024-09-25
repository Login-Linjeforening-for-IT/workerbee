package api

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"gitlab.login.no/tekkom/web/beehive/admin-api/token"

	"github.com/gin-gonic/gin"
)

var (
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrUnauthorized    = errors.New("unauthorized")
)

func (server *Server) authMiddleware(checkAuthorization authorizationCheckFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := getAccessTokenFromHeader(ctx)

		var (
			payload   *token.Payload = nil
			refreshed                = false
		)

		if err != nil {
			payload, err = server.refreshAccessToken(ctx)
			if err == nil {
				refreshed = true
			} else {
				server.writeError(ctx, http.StatusUnauthorized, fmt.Errorf("authMiddleware, refreshAccessToken - %w", ErrUnauthenticated))
				ctx.Abort()
				return
			}
		}

		if !refreshed {
			payload, err = server.accessTokenMaker.VerifyToken(accessToken)
			if err != nil {
				server.writeError(ctx, http.StatusUnauthorized, fmt.Errorf("authMiddleware, VerifyToken - %w", ErrUnauthenticated))
				ctx.Abort()
				return
			}
		}

		if checkAuthorization != nil {
			if !checkAuthorization(payload.Roles) {
				server.writeError(ctx, http.StatusForbidden, fmt.Errorf("authMiddleware, checkAuthorization - %w", ErrUnauthenticated))
				ctx.Abort()
				return
			}
		}

		setUser(ctx, payload)
		ctx.Next()
	}
}

// type authorizationCheckStrategy func(requiredRoles ...string) authorizationCheckFunc
type authorizationCheckFunc func(roles []string) bool

// func allOf(requiredRoles ...string) authorizationCheckFunc {
// 	return func(roles []string) bool {
// 		for _, requiredRole := range requiredRoles {
// 			found := false
// 			for _, role := range roles {
// 				if role == requiredRole {
// 					found = true
// 					break
// 				}
// 			}
// 			if !found {
// 				return false
// 			}
// 		}
// 		return true
// 	}
// }

// func anyOf(requiredRoles ...string) authorizationCheckFunc {
// 	return func(roles []string) bool {
// 		for _, requiredRole := range requiredRoles {
// 			for _, role := range roles {
// 				if role == requiredRole {
// 					return true
// 				}
// 			}
// 		}
// 		return false
// 	}
// }

func regexpMatch(pattern string) authorizationCheckFunc {
	r := regexp.MustCompile(pattern)

	return func(roles []string) bool {
		for _, role := range roles {
			if r.MatchString(role) {
				return true
			}
		}
		return false
	}
}

func (server *Server) refreshAccessToken(ctx *gin.Context) (*token.Payload, error) {
	refreshToken, err := getRefreshTokenFromHeader(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	refreshTokenPayload, err := server.refreshTokenMaker.VerifyToken(refreshToken)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	session, err := server.sessionstore.GetSession(ctx, refreshTokenPayload.ID)
	if err != nil {
		return nil, err
	}

	// validate refresh token against session
	if session.RefreshToken != refreshToken {
		return nil, fmt.Errorf("refresh token not matching: %w", ErrUnauthenticated)
	}

	if session.UID != refreshTokenPayload.UID {
		return nil, fmt.Errorf("user id not matching: %w", ErrUnauthenticated)
	}

	if session.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("refresh token expired: %w", ErrUnauthenticated)
	}

	if !session.BlockedAt.IsZero() {
		return nil, fmt.Errorf("refresh token blocked: %w", ErrUnauthenticated)
	}

	if session.UserAgent != ctx.Request.UserAgent() {
		return nil, fmt.Errorf("user agent not matching: %w", ErrUnauthenticated)
	}

	if session.ClientIP != ctx.ClientIP() {
		return nil, fmt.Errorf("client ip not matching: %w", ErrUnauthenticated)
	}

	accessToken, accessTokenPayload, err := server.createAccessToken(ctx, refreshTokenPayload.UID, refreshTokenPayload.Roles)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}

	setAccessTokenCookie(ctx, accessToken, accessTokenPayload)
	return accessTokenPayload, nil
}

const userContextKey = "user"

func setUser(ctx *gin.Context, payload *token.Payload) {
	ctx.Set(userContextKey, payload)
}

// func getUser(ctx *gin.Context) *token.Payload {
// 	user, ok := ctx.Get(userContextKey)
// 	if !ok {
// 		return nil
// 	}
// 	return user.(*token.Payload)
// }
