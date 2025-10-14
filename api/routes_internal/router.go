package routes_internal

import (
	"workerbee/handlers"
	"workerbee/internal"
	"workerbee/internal/middleware"

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
			events.POST("/", middleware.AuthMiddleware(), h.CreateEvent)
			events.PUT("/:id", middleware.AuthMiddleware(), h.UpdateEvent)
			events.DELETE("/:id", middleware.AuthMiddleware(), h.DeleteEvent)
			events.GET("/categories", h.GetEventCategories)
		}
		rules := v2.Group("/rules")
		{
			rules.GET("/:id", h.GetRule)
			rules.GET("/", h.GetRules)
			rules.POST("/", middleware.AuthMiddleware(), h.CreateRule)
			rules.PUT("/:id", middleware.AuthMiddleware(), h.UpdateRule)
			rules.DELETE("/:id", middleware.AuthMiddleware(), h.DeleteRule)
		}
		locations := v2.Group("/locations")
		{
			locations.GET("/:id", h.GetLocation)
			locations.GET("/", h.GetLocations)
			locations.POST("/", middleware.AuthMiddleware(), h.CreateLocation)
			locations.PUT("/:id", middleware.AuthMiddleware(), h.UpdateLocation)
			locations.DELETE("/:id", middleware.AuthMiddleware(), h.DeleteLocation)
		}
		organizations := v2.Group("/organizations")
		{
			organizations.GET("/:id", h.GetOrganization)
			organizations.GET("/", h.GetOrganizations)
			organizations.POST("/", middleware.AuthMiddleware(), h.CreateOrganization)
			organizations.PUT("/:id", middleware.AuthMiddleware(), h.UpdateOrganization)
			organizations.DELETE("/:id", middleware.AuthMiddleware(), h.DeleteOrganization)
		}
		jobs := v2.Group("/jobs")
		{
			jobs.GET("/:id", h.GetJob)
			jobs.GET("/", h.GetJobs)
			jobs.POST("/", middleware.AuthMiddleware(), h.CreateJob)
			jobs.PUT("/:id", middleware.AuthMiddleware(), h.UpdateJob)
			jobs.DELETE("/:id", middleware.AuthMiddleware(), h.DeleteJob)
			jobs.GET("/cities", h.GetCities)
			jobs.GET("/types", h.GetJobTypes)
			jobs.GET("/skills", h.GetJobSkills)
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
