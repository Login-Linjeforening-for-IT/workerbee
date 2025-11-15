package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

type response struct {
	Sub               string   `json:"sub"`
	Email             string   `json:"email"`
	EmailVerified     bool     `json:"email_verified"`
	Name              string   `json:"name"`
	GivenName         string   `json:"given_name"`
	PreferredUsername string   `json:"preferred_username"`
	Nickname          string   `json:"nickname"`
	Groups            []string `json:"groups"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			internal.HandleError(c, internal.ErrUnauthorized)
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			internal.HandleError(c, internal.ErrUnauthorized)
			c.Abort()
			return
		}

		token := tokenParts[1]

		req, err := http.NewRequest(http.MethodGet, internal.USERINFO_URL, nil)
		if internal.HandleError(c, err) {
			c.Abort()
			return
		}

		req.Header.Set("Authorization", "Bearer "+token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if internal.HandleError(c, err) {
			c.Abort()
			return
		}

		if resp.StatusCode != 200 {
			internal.HandleError(c, internal.ErrUnauthorized)
			c.Abort()
			return
		}

		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		var respStruct response
		if err := decoder.Decode(&respStruct); err != nil {
			internal.HandleError(c, internal.ErrUnauthorized)
			c.Abort()
			return
		}

		if !slices.Contains(respStruct.Groups, internal.QUEENBEE_GROUP) {
			internal.HandleError(c, internal.ErrUnauthorized)
			c.Abort()
			return
		}

		c.Set("user", respStruct.Sub)

		fmt.Printf("[Protected] username=%s\n",
			respStruct.Nickname,
		)

		c.Next()
	}
}
