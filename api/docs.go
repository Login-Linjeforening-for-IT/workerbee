package api

import (
	adminapi "git.logntnu.no/tekkom/web/beehive/admin-api"
	"git.logntnu.no/tekkom/web/beehive/admin-api/docs"
)

func (server *Server) setSwaggerInfo() {
	docs.SwaggerInfo.Title = "Beehive Admin API"
	docs.SwaggerInfo.Description = "Admin API for Beehive"
	docs.SwaggerInfo.Version = adminapi.Version()
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
