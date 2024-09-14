package api

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func AuthentikOauth2Config(baseURL string, clientID string, clientSecret string, redirectURL string) *oauth2Config {
	return &oauth2Config{
		Config: oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       []string{"email", "profile", "openid"},
			Endpoint:     authentikEndpoint(baseURL),
		},
		UserInfoEndpoint: baseURL + OauthAuthentikUserInfoPath,
		RevokeEndpoint:   baseURL + "/application/o/revoke/",
		stateExpiration:  stateExpiration,
		provider:         "authentik",
	}
}

func authentikEndpoint(baseURL string) oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:  baseURL + "/application/o/authorize/",
		TokenURL: baseURL + "/application/o/token/",
	}
}

func (server *Server) authentikCallback() gin.HandlerFunc {
	return server.oauth2Fallback("authentik", server.oauth2Config.getAuthentikUserInfo)
}

var OauthAuthentikUserInfoPath = "/application/o/userinfo/"

type authentikUserInfo struct {
	Sub    string   `json:"sub"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Groups []string `json:"groups"`
}

func (conf *oauth2Config) getAuthentikUserInfo(ctx context.Context, token *oauth2.Token) (userInfo, error) {
	client := conf.Client(ctx, token)

	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	response, err := client.Get(conf.UserInfoEndpoint)
	if err != nil {
		return userInfo{}, fmt.Errorf("userinfo error: %w", err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return userInfo{}, err
	}

	/* e.g.
	{
		"email": "even.testersen@login.no",
		"email_verified": true,
		"name": "Even Testersen",
		"given_name": "Even Testersen",
		"preferred_username": "even-tekkom-verv",
		"nickname": "even-tekkom-verv",
		"groups": ["Tekkom-verv"],
		"sub": "6e30705a8ddc2371c38d0cf767325282645cf6449c9b0b75e1a9a7fec5101312"
	}
	*/
	fmt.Println(string(content))
	
	var u authentikUserInfo
	err = json.Unmarshal(content, &u)
	// err = json.NewDecoder(response.Body).Decode(&u)
	if err != nil {
		return userInfo{}, fmt.Errorf("getAuthentikUserInfo - Unmarshal failed: %s\nJSON content: %s, Length: %d", err.Error(), string(content), len(content))
	}

	return userInfo{
		ID:    u.Sub,
		Name:  u.Name,
		Email: u.Email,
		Roles: u.Groups,
	}, nil
}
