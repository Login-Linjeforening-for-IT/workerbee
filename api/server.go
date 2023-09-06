package api

import (
	"os"

	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
	"git.logntnu.no/tekkom/web/beehive/admin-api/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

func init() {
	validator, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}

	db.BindTimeTypeEnumValidator(validator, "timetypeenum")
	db.BindLocationTypeValidator(validator, "locationtype")
}

type Config struct {
	Port string `config:"PORT" default:"8080"`
}

type Server struct {
	config  *Config
	router  *gin.Engine
	service service.Service
	logger  zerolog.Logger
}

func NewServer(config *Config, service service.Service) *Server {
	server := &Server{
		config:  config,
		service: service,
		logger:  zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}

	server.initRouter()

	return server
}

func (server *Server) initRouter() {
	router := gin.Default()

	api := router.Group("/api", server.authMiddleware)
	{
		events := api.Group("/events")
		{
			events.GET("/", server.getEvents)
			events.GET("/:id", server.getEvent)
			events.POST("/", server.createEvent)
			events.PATCH("/", server.updateEvent)

			// rules := events.Group("/rules")
			// {
			// 	rules.POST("/", server.addRuleToEvent)
			// 	rules.DELETE("/", server.removeRuleFromEvent)
			// }
		}

		rules := api.Group("/rules")
		{
			rules.GET("/", server.getRules)
			rules.GET("/:id", server.getRule)
			rules.POST("/", server.createRule)
			rules.PATCH("/", server.updateRule)
		}

		locations := api.Group("/locations")
		{
			locations.GET("/", server.getLocations)
			locations.GET("/:id", server.getLocation)
			locations.POST("/", server.createLocation)
			locations.PATCH("/", server.updateLocation)
		}

		organizations := api.Group("/organizations")
		{
			organizations.GET("/", server.getOrganizations)
			organizations.GET("/:shortname", server.getOrganization)
			organizations.POST("/", server.createOrganization)
			organizations.PATCH("/", server.updateOrganization)
		}

		categories := api.Group("/categories")
		{
			categories.GET("/", server.getCategories)
			categories.GET("/:id", server.getCategory)
		}
	}

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(":" + server.config.Port)
}
