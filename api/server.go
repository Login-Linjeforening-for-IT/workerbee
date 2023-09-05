package api

import (
	"os"

	"git.logntnu.no/tekkom/web/beehive/admin-api/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

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
		}
	}

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(":" + server.config.Port)
}
