package internal

import (
	"github.com/gin-gonic/gin"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/handlers"
)

func Route(c *gin.Engine) {
	v1 := c.Group(BASE_PATH)
	{
		v1.GET("/ping", handlers.PongHandler)
		events := v1.Group("/events")
		{
			events.GET("/testing", handlers.PongHandler)
		}
		rules := v1.Group("/rules")
		{
			rules.GET("/testing", handlers.PongHandler)
		}
		locations := v1.Group("/locations")
		{
			locations.GET("/testing", handlers.PongHandler)
		}
		organizations := v1.Group("/organizations")
		{
			organizations.GET("/testing", handlers.PongHandler)
		}
		categories := v1.Group("/categories")
		{
			categories.GET("/testing", handlers.PongHandler)
		}
		audiences := v1.Group("/audiences")
		{
			audiences.GET("/testing", handlers.PongHandler)
		}
		jobs := v1.Group("/jobs")
		{
			jobs.GET("/testing", handlers.PongHandler)
		}
		cities := v1.Group("/cities")
		{
			cities.GET("/testing", handlers.PongHandler)
		}

	}
}
