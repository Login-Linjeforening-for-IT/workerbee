package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type request struct {
	Bearer string `json:"bearer"`
	Token  string `json:"token"`
}

func AuthMiddelware(audience string, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token := tokenParts[1]

		rawData := request{Bearer: "Bearer", Token: token}
		jsonData, err := json.Marshal(rawData)
		if err != nil {
			log.Println("Failed to marshal token for authentication")
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		req, err := http.NewRequest(http.MethodPost, USERINGO_URL, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("Unable create request to authentik err: ", err)
		}

		req.Header.Set("Content-Type", "Authorization")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Unable to connect to authentik endpoint err: ", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Unable to read body")
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		log.Println(body)

		c.Next()

	}
}
