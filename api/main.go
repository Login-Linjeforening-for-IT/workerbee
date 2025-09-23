package main

import (
	"workerbee/config"
	"workerbee/db"
	"workerbee/handlers"
	repositories "workerbee/repositories"
	"workerbee/routes_internal"
	"workerbee/services"

	"github.com/gin-gonic/gin"
)

func init() {
	/*
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	*/
	config.Init()
}

func main() {
	db := db.Init()

	// Repos
	repos := repositories.NewRepositories(db)

	// Services
	svcs := services.NewServices(repos)

	// handler container
	h := &handlers.Handler{ Services: svcs }

	router := gin.New()

	router.Use(gin.Logger())

	routes_internal.Route(router, h)

	router.Run(":" + config.Port)
}
