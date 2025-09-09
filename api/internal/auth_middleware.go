package internal

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"
	"strings"

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
			HandleError(c, nil, "No Bearer registerd in headers", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			HandleError(c, nil, "Unable to extract token bad format", http.StatusBadRequest)
			return
		}

		token := tokenParts[1]

		req, err := http.NewRequest(http.MethodGet, USERINFO_URL, nil)
		if err != nil {
			HandleError(c, err, "Unable to create request to Authentik", http.StatusInternalServerError)
			return
		}

		req.Header.Set("Authorization", "Bearer "+token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			HandleError(c, err, "Unable to make request to Authentik", http.StatusInternalServerError)
			return
		}

		if resp.StatusCode != 200 {
			HandleError(c, nil, "Expected "+strconv.Itoa(http.StatusOK)+" but got "+strconv.Itoa(resp.StatusCode), http.StatusUnauthorized)
			return
		}

		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		var respStruct response
		if err := decoder.Decode(&respStruct); err != nil {
			HandleError(c, err, "Token not registerd unable to decode response", http.StatusUnauthorized)
			return
		}

		if !slices.Contains(respStruct.Groups, QUEENBEE_GROUP) {
			HandleError(c, err, "User not in group "+QUEENBEE_GROUP, http.StatusUnauthorized)
			return
		}

		c.Next()

	}
}
