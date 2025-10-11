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
			events.POST("/", internal.AuthMiddleware(), h.CreateEvent)
			events.PUT("/:id", internal.AuthMiddleware(), h.UpdateEvent)
			events.DELETE("/:id", internal.AuthMiddleware(), h.DeleteEvent)
			events.GET("/categories", h.GetCategories)
		}
		rules := v2.Group("/rules")
		{
			rules.GET("/:id", h.GetRule)
			rules.GET("/", h.GetRules)
			rules.POST("/", internal.AuthMiddleware(), h.CreateRule)
			rules.PUT("/:id", internal.AuthMiddleware(), h.UpdateRule)
			rules.DELETE("/:id", internal.AuthMiddleware(), h.DeleteRule)
		}
		locations := v2.Group("/locations")
		{
			locations.GET("/:id", h.GetLocation)
			locations.GET("/", h.GetLocations)
			locations.POST("/", internal.AuthMiddleware(), h.CreateLocation)
			locations.PUT("/:id", internal.AuthMiddleware(), h.UpdateLocation)
			locations.DELETE("/:id", internal.AuthMiddleware(), h.DeleteLocation)
		}
		organizations := v2.Group("/organizations")
		{
			organizations.GET("/:id", h.GetOrganization)
			organizations.GET("/", h.GetOrganizations)
			organizations.POST("/", internal.AuthMiddleware(), h.CreateOrganization)
			organizations.PUT("/:id", internal.AuthMiddleware(), h.UpdateOrganization)
			organizations.DELETE("/:id", internal.AuthMiddleware(), h.DeleteOrganization)
		}
		categories := v2.Group("/categories")
		{
			categories.GET("/:id", h.GetCategory)
			categories.GET("/", h.GetCategories)
			categories.POST("/", internal.AuthMiddleware(), h.CreateCategory)
			categories.PUT("/:id", internal.AuthMiddleware(), h.UpdateCategory)
			categories.DELETE("/:id", internal.AuthMiddleware(), h.DeleteCategory)
		}
		jobs := v2.Group("/jobs")
		{
			jobs.GET("/:id", h.GetJob)
			jobs.GET("/", h.GetJobs)
			jobs.POST("/", internal.AuthMiddleware(), h.CreateJob)
			jobs.PUT("/:id", internal.AuthMiddleware(), h.UpdateJob)
			jobs.DELETE("/:id", internal.AuthMiddleware(), h.DeleteJob)
			jobs.GET("/cities", h.GetCities)
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
