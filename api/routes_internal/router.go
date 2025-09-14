package routes_internal

import (
	"github.com/gin-gonic/gin"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/handlers"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/internal"
)

func Route(c *gin.Engine) {
	v2 := c.Group(internal.BASE_PATH)
	{
		v2.GET("/ping", handlers.PingHandler)
		v2.GET("/docs", handlers.GetDocs)
		v2.GET("/status", handlers.GetStatus)
		events := v2.Group("/events")
		{
			events.GET("/:id", handlers.PingHandler)
			events.GET("/", handlers.GetEvents)
			events.POST("/", handlers.PingHandler)
			events.PATCH("/:id", handlers.PingHandler)
			events.DELETE("/:id", handlers.PingHandler)
			events.GET("/categories", handlers.PingHandler)
			events.GET("/audience", handlers.PingHandler)
		}
		rules := v2.Group("/rules")
		{
			rules.GET("/:id", handlers.PingHandler)
			rules.GET("/", handlers.PingHandler)
			rules.POST("/", handlers.PingHandler)
			rules.DELETE("/:id", handlers.PingHandler)
		}
		locations := v2.Group("/locations")
		{
			locations.GET("/:id", handlers.PingHandler)
			locations.GET("/", handlers.PingHandler)
			locations.POST("/", handlers.PingHandler)
			locations.PATCH("/:id", handlers.PingHandler)
			locations.DELETE("/:id", handlers.PingHandler)
		}
		organizations := v2.Group("/organizations")
		{
			organizations.GET("/:id", handlers.PingHandler)
			organizations.GET("/", handlers.PingHandler)
			organizations.POST("/", handlers.PingHandler)
			organizations.DELETE("/:id", handlers.PingHandler)
			organizations.PATCH("/:id", handlers.PingHandler)
		}
		categories := v2.Group("/categories")
		{
			categories.GET("/", handlers.PingHandler)
			categories.GET("/:id", handlers.PingHandler)
			categories.POST("/", handlers.PingHandler)

		}
		jobs := v2.Group("/jobs")
		{
			jobs.GET("/:id", handlers.PingHandler)
			jobs.GET("/", handlers.PingHandler)
			jobs.POST("/", handlers.PingHandler)
			jobs.PATCH("/:id", handlers.PingHandler)
			jobs.DELETE("/:id", handlers.PingHandler)
			jobs.GET("/cities", handlers.PingHandler)
		}
		stats := v2.Group("/stats")
		{
			stats.GET("/total", handlers.GetTotalStats)
			stats.GET("/categories", handlers.GetCategoriesStats)
			stats.GET("/new-additions", handlers.GetNewAdditionsStats)
		}
		forms := v2.Group("/forms")
		{
			forms.GET("/:id", handlers.GetForm)
			forms.GET("/", handlers.GetForms)
			forms.POST("/", handlers.PingHandler)
			forms.DELETE("/:id", handlers.PingHandler)
			forms.PATCH("/:id", handlers.PingHandler)
			questions := forms.Group("/:id/questions")
			{
				questions.GET("/", handlers.PingHandler)
				questions.POST("/", handlers.PingHandler)
				questions.DELETE("/", handlers.PingHandler)
				questions.PATCH("/", handlers.PingHandler)
			}
			answers := forms.Group("/:id/answers")
			{
				answers.GET("/", handlers.PingHandler)
				answers.POST("/", handlers.PingHandler)
				answers.DELETE("/:id", handlers.PingHandler)
				answers.PATCH("/:id", handlers.PingHandler)
			}
		}
	}
}
