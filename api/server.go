package api

import (
	"net/http"
	"os"

	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
	"git.logntnu.no/tekkom/web/beehive/admin-api/service"
	"github.com/gin-contrib/cors"
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
	db.BindJobTypeValidator(validator, "jobtype")
}

type Config struct {
	Port   string `config:"PORT" default:"8080"`
	Secret string `config:"SECRET" default:"secret"`

	AllowedHeaders []string `config:"ALLOWED_HEADERS" default:"Content-Type,Authorization,Accept,Accept-Encoding,Accept-Language"`
	AllowedMethods []string `config:"ALLOWED_METHODS" default:"GET,POST,PUT,PATCH,DELETE,OPTIONS"`
	AllowedOrigins []string `config:"ALLOWED_ORIGINS" default:"*"`

	DOKey    string `config:"DO_ACCESS_KEY_ID"`
	DOSecret string `config:"DO_SECRET_ACCESS_KEY"`
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

	corsConf := cors.DefaultConfig()
	if server.config.AllowedOrigins != nil && len(server.config.AllowedOrigins) > 0 {
		corsConf.AllowOrigins = server.config.AllowedOrigins
	} else {
		corsConf.AllowAllOrigins = true
	}
	corsConf.AllowCredentials = true
	corsConf.AllowHeaders = server.config.AllowedHeaders
	corsConf.AllowMethods = server.config.AllowedMethods

	router.Use(cors.New(corsConf))

	router.NoRoute(server.noRoute)

	api := router.Group("/api", server.authMiddleware(server.config.Secret))
	{
		events := api.Group("/events")
		{
			events.GET("/", server.getEvents)
			events.GET("/:id", server.getEvent)
			events.POST("/", server.createEvent)
			events.PATCH("/", server.updateEvent)
			events.DELETE("/:id", server.deleteEvent)

			events.POST("/organizations", server.addOrganizationToEvent)
			events.DELETE("/organizations", server.removeOrganizationFromEvent)

			events.POST("/audiences", server.addAudienceToEvent)
			events.DELETE("/audiences", server.removeAudienceFromEvent)
		}

		rules := api.Group("/rules")
		{
			rules.GET("/", server.getRules)
			rules.GET("/:id", server.getRule)
			rules.POST("/", server.createRule)
			rules.PATCH("/", server.updateRule)
			rules.DELETE("/:id", server.deleteRule)
		}

		locations := api.Group("/locations")
		{
			locations.GET("/", server.getLocations)
			locations.GET("/:id", server.getLocation)
			locations.POST("/", server.createLocation)
			locations.PATCH("/", server.updateLocation)
			locations.DELETE("/:id", server.deleteLocation)
		}

		organizations := api.Group("/organizations")
		{
			organizations.GET("/", server.getOrganizations)
			organizations.GET("/:shortname", server.getOrganization)
			organizations.POST("/", server.createOrganization)
			organizations.PATCH("/", server.updateOrganization)
			organizations.DELETE("/:shortname", server.deleteOrganization)
		}

		categories := api.Group("/categories")
		{
			categories.GET("/", server.getCategories)
			categories.GET("/:id", server.getCategory)
		}

		audiences := api.Group("/audiences")
		{
			audiences.GET("/", server.getAudiences)
			audiences.GET("/:id", server.getAudience)
		}

		jobs := api.Group("/jobs")
		{
			jobs.GET("/", server.getJobs)
			jobs.GET("/:id", server.getJob)
			jobs.POST("/", server.createJob)
			jobs.PATCH("/", server.updateJob)
			jobs.DELETE("/:id", server.deleteJob)

			jobs.POST("/cities", server.addCityToJob)
			jobs.DELETE("/cities", server.removeCityFromJob)

			jobs.POST("/skills", server.addSkillToJob)
			jobs.DELETE("/skills", server.removeSkillFromJob)
		}

		cities := api.Group("/cities")
		{
			cities.GET("/", server.getAllCities)
		}

		images := api.Group("/images")
		{
			images.GET("/events/banner", server.fetchEventsBannerList)

			images.POST("/events/banner", server.uploadEventImageBanner)
			images.POST("/events/small", server.uploadEventImageSmall)
			images.POST("/jobs", server.uploadAdImage)
			images.POST("/organizations", server.uploadOrganizationImage)
		}
	}

	server.router = router
}

func (server *Server) noRoute(ctx *gin.Context) {
	server.writeError(ctx, http.StatusNotFound, &NotFoundError{
		Message: ctx.Request.URL.Path + " not found",
	})
}

func (server *Server) Start() error {
	return server.router.Run(":" + server.config.Port)
}
