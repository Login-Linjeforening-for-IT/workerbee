package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/internal"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/internal/config"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/internal/db"
)

func init() {
	/*
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	*/
	config.Init()
	db.Init()
}

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	internal.Route(router)

	router.Run(":" + config.Port)
}
