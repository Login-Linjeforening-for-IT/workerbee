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
	eventRepo := repositories.NewEventrepositories(db)
	ruleRepo := repositories.NewRulerepositories(db)
	jobRepo := repositories.NewJobrepositories(db)
	statsRepo := repositories.NewStatsrepositories(db)
	formRepo := repositories.NewFormrepositories(db)
	questionRepo := repositories.NewQuestionrepositories(db)
	submissionRepo := repositories.NewSubmissionrepositories(db)

	// Services
	eventService := services.NewEventService(eventRepo)
	ruleService := services.NewRuleService(ruleRepo)
	jobService := services.NewJobsService(jobRepo)
	statsService := services.NewStatsService(statsRepo)
	formService := services.NewFormService(formRepo)
	questionService := services.NewQuestionService(questionRepo)
	submissionService := services.NewSubmissionService(submissionRepo)

	// handler container
	h := &handlers.Handler{
		Events:      *eventService,
		Rules:       *ruleService,
		Jobs:        *jobService,
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
