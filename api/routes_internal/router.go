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
			events.GET("/protected/:id", middleware.AuthMiddleware(), h.GetProtectedEvent)
			events.GET("/:id", h.GetEvent)
			events.GET("/protected", middleware.AuthMiddleware(), h.GetProtectedEvents)
			events.GET("/", h.GetEvents)
			events.POST("/", middleware.AuthMiddleware(), h.CreateEvent)
			events.PUT("/:id", middleware.AuthMiddleware(), h.UpdateEvent)
			events.DELETE("/:id", middleware.AuthMiddleware(), h.DeleteEvent)
			events.GET("/categories", h.GetEventCategories)
			events.GET("/audiences", h.GetEventAudiences)
			events.GET("/time", h.GetAllTimeTypes)
		}
		rules := v2.Group("/rules")
		{
			rules.GET("/:id", h.GetRule)
			rules.GET("/all", h.GetRuleNames)
			rules.GET("/", h.GetRules)
			rules.POST("/", middleware.AuthMiddleware(), h.CreateRule)
			rules.PUT("/:id", middleware.AuthMiddleware(), h.UpdateRule)
			rules.DELETE("/:id", middleware.AuthMiddleware(), h.DeleteRule)
		}
		categories := v2.Group("/categories")
		{
			categories.GET("/:id", h.GetCategory)
			categories.GET("/", h.GetCategories)
			categories.POST("/", middleware.AuthMiddleware(), h.CreateCategory)
			categories.PUT("/:id", middleware.AuthMiddleware(), h.UpdateCategory)
			categories.DELETE("/:id", middleware.AuthMiddleware(), h.DeleteCategory)
		}
		locations := v2.Group("/locations")
		{
			locations.GET("/:id", h.GetLocation)
			locations.GET("/all", h.GetLocationNames)
			locations.GET("/", h.GetLocations)
			locations.POST("/", middleware.AuthMiddleware(), h.CreateLocation)
			locations.PUT("/:id", middleware.AuthMiddleware(), h.UpdateLocation)
			locations.DELETE("/:id", middleware.AuthMiddleware(), h.DeleteLocation)
			locations.GET("/types", h.GetAllLocationTypes)
		}
		organizations := v2.Group("/organizations")
		{
			organizations.GET("/:id", h.GetOrganization)
			organizations.GET("/all", h.GetOrganizationNames)
			organizations.GET("/", h.GetOrganizations)
			organizations.POST("/", middleware.AuthMiddleware(), h.CreateOrganization)
			organizations.PUT("/:id", middleware.AuthMiddleware(), h.UpdateOrganization)
			organizations.DELETE("/:id", middleware.AuthMiddleware(), h.DeleteOrganization)
		}
		jobs := v2.Group("/jobs")
		{
			jobs.GET("/:id", h.GetJob)
			jobs.GET("/protected/:id", middleware.AuthMiddleware(), h.GetProtectedJob)
			jobs.GET("/", h.GetJobs)
			jobs.GET("/protected", middleware.AuthMiddleware(), h.GetProtectedJobs)
			jobs.POST("/", h.CreateJob)
			jobs.PUT("/:id", middleware.AuthMiddleware(), h.UpdateJob)
			jobs.DELETE("/:id", middleware.AuthMiddleware(), h.DeleteJob)
			jobs.GET("/cities", h.GetCities)
			jobs.GET("/skills", h.GetJobSkills)
			types := jobs.Group("/types")
			{
				types.GET("/", h.GetActiveJobTypes)
				types.GET("/all", h.GetAllJobTypes)
				types.GET("/:id", h.GetJobType)
				types.POST("/", middleware.AuthMiddleware(), h.CreateJobType)
				types.PUT("/:id", middleware.AuthMiddleware(), h.UpdateJobType)
				types.DELETE("/:id", middleware.AuthMiddleware(), h.DeleteJobType)
			}

		}
		audiences := v2.Group("/audiences")
		{
			audiences.GET("/:id", h.GetAudience)
			audiences.GET("/", h.GetAudiences)
			audiences.POST("/", middleware.AuthMiddleware(), h.CreateAudience)
			audiences.PUT("/:id", middleware.AuthMiddleware(), h.UpdateAudience)
			audiences.DELETE("/:id", middleware.AuthMiddleware(), h.DeleteAudience)
		}
		stats := v2.Group("/stats")
		{
			stats.GET("/yearly", h.GetYearlyStats)
			stats.GET("/category", h.GetMostActiveCategory)
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
		images := v2.Group("/images/:path")
		{
			images.POST("", middleware.AuthMiddleware(), h.UploadImage)
			images.GET("/", middleware.AuthMiddleware(), h.GetImageURLs)
			images.DELETE("/:imageName", middleware.AuthMiddleware(), h.DeleteImage)
		}
		text := v2.Group("/text")
		{
			text.GET("/", h.GetTextServices)
			service := text.Group("/:service")
			{
				service.GET("/", h.GetAllPathsInService)
				content := service.Group("/:path")
				{
					content.GET("/", h.GetAllContentInPath)
					content.PUT("/", middleware.AuthMiddleware(), h.UpdateContentInPath)
					content.POST("/", middleware.AuthMiddleware(),h.CreateTextInService)
					language := content.Group("/:language")
					{
						language.GET("/", h.GetOneLanguage)
					}
				}
			}
		}
	}
}
