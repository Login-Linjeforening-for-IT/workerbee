package internal

import (
	"github.com/gin-gonic/gin"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/handlers"
)

func Route(c *gin.Engine) {
	v1 := c.Group(BASE_PATH)
	{
		v1.GET("/ping", AuthMiddleware(), handlers.PongHandler)
		events := v1.Group("/events")
		{
			events.GET("/:id", handlers.PongHandler)
			events.GET("/", handlers.PongHandler)
			events.POST("/", handlers.PongHandler)
			events.PATCH("/:id", handlers.PongHandler)
			events.DELETE("/:id", handlers.PongHandler)
			events.GET("/categories", handlers.PongHandler)
			events.GET("/audience", handlers.PongHandler)
		}
		rules := v1.Group("/rules")
		{
			rules.GET("/:id", handlers.PongHandler)
			rules.GET("/", handlers.PongHandler)
			rules.POST("/", handlers.PongHandler)
			rules.DELETE("/:id", handlers.PongHandler)
		}
		locations := v1.Group("/locations")
		{
			locations.GET("/:id", handlers.PongHandler)
			locations.GET("/", handlers.PongHandler)
			locations.POST("/", handlers.PongHandler)
			locations.PATCH("/:id", handlers.PongHandler)
			locations.DELETE("/:id", handlers.PongHandler)
		}
		organizations := v1.Group("/organizations")
		{
			organizations.GET("/:id", handlers.PongHandler)
			organizations.GET("/", handlers.PongHandler)
			organizations.POST("/", handlers.PongHandler)
			organizations.DELETE("/:id", handlers.PongHandler)
			organizations.PATCH("/:id", handlers.PongHandler)
		}
		categories := v1.Group("/categories")
		{
			categories.GET("/", handlers.PongHandler)
			categories.GET("/:id", handlers.PongHandler)
			categories.POST("/", handlers.PongHandler)

		}
		jobs := v1.Group("/jobs")
		{
			jobs.GET("/:id", handlers.PongHandler)
			jobs.GET("/", handlers.PongHandler)
			jobs.POST("/", handlers.PongHandler)
			jobs.PATCH("/:id", handlers.PongHandler)
			jobs.DELETE("/:id", handlers.PongHandler)
			jobs.GET("/cities", handlers.PongHandler)
		}
	}
}
