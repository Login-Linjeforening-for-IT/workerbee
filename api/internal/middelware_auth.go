package internal

/*
type authentikResponse struct {
	Active bool `json:"active"`
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

		token := tokenParts[1]

		data := url.Values{}
		data.Set("token", token)
		data.Set("token_type_hint", "access_token") // optional but recommended

		req, _ := http.NewRequest("POST", config.AuthentikURL, strings.NewReader(data.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// Optional if your provider allows unauthenticated introspection
		req.SetBasicAuth(audience, secret)

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// Print the response as string
		fmt.Println(string(body))

		c.Next()
	}
}
*/
