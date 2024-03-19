package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"golang.org/x/oauth2"
)

type userInfo struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
	// Picture string `json:"picture"`
}

type oauth2Config struct {
	oauth2.Config
	UserInfoEndpoint string
	RevokeEndpoint   string
	stateExpiration  time.Duration
	provider         string
}

var stateExpiration = 20 * time.Minute

func (conf *oauth2Config) generateStateOauthCookie(ctx *gin.Context) string {
	// expiration := time.Now().Add(conf.stateExpiration)
	state := generateOauthState()

	// cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	// http.SetCookie(w, &cookie)
	ctx.SetCookie("oauthstate", state, int(conf.stateExpiration.Seconds()), "/", "", false, true)

	return state
}

func generateOauthState() string {
	return xid.New().String()
}

// exchange code for token
func (conf oauth2Config) getToken(code string) (*oauth2.Token, error) {
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %w", err)
	}

	return token, nil
}

func (server *Server) oauth2Login(ctx *gin.Context) {
	oauthState := server.oauth2Config.generateStateOauthCookie(ctx)
	u := server.oauth2Config.AuthCodeURL(oauthState)
	ctx.Redirect(http.StatusTemporaryRedirect, u)
}

type getUserInfoFunc func(ctx context.Context, token *oauth2.Token) (userInfo, error)

func (server *Server) oauth2Fallback(provider string, getUserInfo getUserInfoFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		oauthState, err := ctx.Request.Cookie("oauthstate")
		if err != nil || ctx.Request.FormValue("state") != oauthState.Value {
			err = fmt.Errorf("error - " + ctx.Request.FormValue("state"))

			server.writeError(ctx, http.StatusUnauthorized, err)
			return
		}

		code := ctx.Request.FormValue("code")
		userToken, err := server.oauth2Config.getToken(code)
		if err != nil {
			server.writeError(ctx, http.StatusInternalServerError, err)
			return
		}

		userInfo, err := getUserInfo(ctx, userToken)
		if err != nil {
			server.writeError(ctx, http.StatusInternalServerError, err)
			return
		}

		// TODO(session-store): Find a solution for this
		// user, err := server.service.UpsertUser(ctx, db.UpsertUserParams{
		// 	Provider:       provider,
		// 	ProviderUserID: userInfo.ID,
		// 	Name:           userInfo.Name,
		// 	Email:          userInfo.Email,
		// })
		// if err != nil {
		// 	server.writeDBError(ctx, err)
		// 	return
		// }

		// TODO(session-store): swap userInfo for user when session store is implemented
		accessToken, accessTokenPayload, err := server.createAccessToken(ctx, userInfo.ID, userInfo.Roles)
		if err != nil {
			server.writeError(ctx, http.StatusInternalServerError, err)
			return
		}

		// TODO(session-store): swap userInfo for user when session store is implemented
		refreshToken, refreshTokenPayload, err := server.createRefreshToken(ctx, userInfo.ID, userInfo.Roles)
		if err != nil {
			server.writeError(ctx, http.StatusInternalServerError, err)
			return
		}

		// TODO(session-store): Find a solution for this
		// err = server.service.CreateSession(ctx, db.CreateSessionParams{
		// 	ID:           refreshTokenPayload.ID,
		// 	UserID:       user.ID,
		// 	RefreshToken: refreshToken,
		// 	UserAgent:    ctx.Request.UserAgent(),
		// 	ClientIp:     ctx.ClientIP(),
		// 	ExpiresAt:    refreshTokenPayload.ExpiresAt,
		// })
		// if err != nil {
		// 	server.writeDBError(ctx, err)
		// 	return
		// }

		setAccessTokenCookie(ctx, accessToken, accessTokenPayload)
		setRefreshTokenCookie(ctx, refreshToken, refreshTokenPayload)

		// ctx.Redirect(http.StatusTemporaryRedirect, "/users/me")

		// TODO(session-store): swap userInfo for user when session store is implemented
		ctx.JSON(http.StatusOK, gin.H{
			"access_token":             accessToken,
			"access_token_expires_at":  accessTokenPayload.ExpiresAt,
			"refresh_token":            refreshToken,
			"refresh_token_expires_at": refreshTokenPayload.ExpiresAt,
			"user": gin.H{
				"id":    userInfo.ID,
				"name":  userInfo.Name,
				"roles": accessTokenPayload.Roles,
			},
		})
	}
}
