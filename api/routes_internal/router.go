package routes_internal

import (
	"workerbee/handlers"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

func Route(c *gin.Engine, h *handlers.Handler) {
	v2 := c.Group(internal.BASE_PATH)
	{
		v2.GET("/ping", handlers.PingHandler)
		v2.GET("/docs", handlers.GetDocs)
		v2.GET("/status", handlers.GetStatus)
		events := v2.Group("/events")
		{
			events.GET("/:id", h.GetEvent)
			events.GET("/", h.GetEvents)
			events.POST("/", handlers.PingHandler)
			events.PUT("/:id", handlers.PingHandler)
			events.DELETE("/:id", h.DeleteEvent)
			events.GET("/categories", handlers.PingHandler)
		}
		rules := v2.Group("/rules")
		{
			rules.GET("/:id", handlers.PingHandler)
			rules.GET("/", handlers.PingHandler)
			rules.POST("/", handlers.PingHandler)
			rules.PUT("/:id", handlers.PingHandler)
			rules.DELETE("/:id", handlers.PingHandler)
		}
		locations := v2.Group("/locations")
		{
			locations.GET("/:id", handlers.PingHandler)
			locations.GET("/", handlers.PingHandler)
			locations.POST("/", handlers.PingHandler)
			locations.PUT("/:id", handlers.PingHandler)
			locations.DELETE("/:id", handlers.PingHandler)
		}
		organizations := v2.Group("/organizations")
		{
			organizations.GET("/:id", handlers.PingHandler)
			organizations.GET("/", handlers.PingHandler)
			organizations.POST("/", handlers.PingHandler)
			organizations.DELETE("/:id", handlers.PingHandler)
			organizations.PUT("/:id", handlers.PingHandler)
		}
		categories := v2.Group("/categories")
		{
			categories.GET("/", handlers.PingHandler)
			categories.GET("/:id", handlers.PingHandler)
			categories.POST("/", handlers.PingHandler)
			categories.PUT("/:id", handlers.PingHandler)
			categories.DELETE("/:id", handlers.PingHandler)
		}
		jobs := v2.Group("/jobs")
		{
			jobs.GET("/:id", h.GetJob)
			jobs.GET("/", h.GetJobs)
			jobs.POST("/", handlers.PingHandler)
			jobs.PUT("/:id", handlers.PingHandler)
			jobs.DELETE("/:id", h.DeleteJob)
			jobs.GET("/cities", handlers.PingHandler)
		}
		stats := v2.Group("/stats")
		{
			stats.GET("/total", h.GetTotalStats)
			stats.GET("/categories", h.GetCategoriesStats)
			stats.GET("/new-additions", h.GetNewAdditionsStats)
		}
		forms := v2.Group("/forms")
		{
			forms.GET("/:id", h.GetForm)
			forms.GET("/", h.GetForms)
			forms.POST("/", h.PostForm)
			forms.PUT("/:id", h.PutForm)
			forms.DELETE("/:id", h.DeleteForm)
			submissions := forms.Group(":id/submissions")
			{
				submissions.GET("/:submission_id", h.GetSubmission)
				submissions.GET("/", handlers.PingHandler)
				submissions.POST("/", handlers.PingHandler)
				submissions.PUT("/:submission_id", handlers.PingHandler)
				submissions.DELETE("/:submission_id", handlers.PingHandler)
			}
		}
	}
}
