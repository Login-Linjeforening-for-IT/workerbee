package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/config"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/internal"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Init()
}

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	internal.Route(router)

	router.Run(":" + config.Port)
}
