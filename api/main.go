package main

import (
	"workerbee/config"
	client "workerbee/db"
	"workerbee/handlers"
	"workerbee/internal/middleware"
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
	db := client.Init()
	do := client.DOInit()

	// Repos
	repos := repositories.NewRepositories(db, do)

	// Services
	svcs := services.NewServices(repos)

	router := gin.New()

	// handler container
	h := &handlers.Handler{
		Services: svcs,
		Router:   router,
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	routes_internal.Route(router, h)

	router.Run(":" + config.Port)
}
