package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/config"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/db"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/handlers"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/repository"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/routes_internal"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/services"
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
	eventRepo := repository.NewEventRepository(db)
	statsRepo := repository.NewStatsRepository(db)
	formRepo := repository.NewFormRepository(db)
	questionRepo := repository.NewQuestionRepository(db)
	submissionRepo := repository.NewSubmissionRepository(db)

	// Services
	eventService := services.NewEventService(eventRepo)
	statsService := services.NewStatsService(statsRepo)
	formService := services.NewFormService(formRepo)
	questionService := services.NewQuestionService(questionRepo)
	submissionService := services.NewSubmissionService(submissionRepo)

	// handler container
	h := &handlers.Handler{
		Events:      *eventService,
		Stats:       *statsService,
		Forms:       *formService,
		Questions:   *questionService,
		Submissions: *submissionService,
	}

	router := gin.New()

	router.Use(gin.Logger())

	routes_internal.Route(router, h)

	router.Run(":" + config.Port)
}
