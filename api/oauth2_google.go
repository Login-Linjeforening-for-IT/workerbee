package api

// "golang.org/x/oauth2/google"

// func GoogleOauth2Config(clientID string, clientSecret string, redirectURL string) *oauth2Config {
// 	return &oauth2Config{
// 		Config: oauth2.Config{
// 			ClientID:     clientID,
// 			ClientSecret: clientSecret,
// 			RedirectURL:  redirectURL,
// 			Scopes: []string{
// 				"https://www.googleapis.com/auth/userinfo.email",
// 				"https://www.googleapis.com/auth/userinfo.profile",
// 			},
// 			Endpoint: google.Endpoint,
// 		},
// 		UserInfoEndpoint: OauthGoogleUserInfoEndpoint,
// 		RevokeEndpoint:   "https://accounts.google.com/o/oauth2/revoke", // check
// 		stateExpiration:  stateExpiration,
// 		provider:         "google",
// 	}
// }

// func (server *Server) googleCallback() gin.HandlerFunc {
// 	return server.oauth2Fallback("google", server.oauth2Config.getGoogleUserInfo)
// }

// const OauthGoogleUserInfoEndpoint = "https://www.googleapis.com/oauth2/v2/userinfo"

// func (conf *oauth2Config) getGoogleUserInfo(ctx context.Context, token *oauth2.Token) (userInfo, error) {
// 	client := conf.Client(ctx, token)
// 	response, err := client.Get(conf.UserInfoEndpoint)
// 	if err != nil {
// 		return userInfo{}, err
// 	}
// 	defer response.Body.Close()

// 	var u userInfo
// 	err = json.NewDecoder(response.Body).Decode(&u)
// 	if err != nil {
// 		return userInfo{}, err
// 	}

// 	return u, nil
// }
