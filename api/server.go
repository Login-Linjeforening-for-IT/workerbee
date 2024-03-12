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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	AllowedHeaders []string `config:"ALLOWED_HEADERS" defult:"Content-Type,Authorization,Accept,Accept-Encoding,Accept-Language"`
	AllowedMethods []string `config:"ALLOWED_METHODS" defult:"GET,POST,PUT,PATCH,DELETE,OPTIONS"`
	AllowedOrigins []string `config:"ALLOWED_ORIGINS" defult:"*"`
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

	server.setSwaggerInfo()

	server.initRouter()

	return server
}

func (server *Server) initRouter() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.CustomRecoveryWithWriter(nil, server.CustomRecovery()))

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
	}

	router.GET("/docs", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(":" + server.config.Port)
}

func (server *Server) StartTLS(certFile, keyFile string) error {
	return server.router.RunTLS(":"+server.config.Port, certFile, keyFile)
}
