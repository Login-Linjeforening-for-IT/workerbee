package api

import (
	adminapi "gitlab.login.no/tekkom/web/beehive/admin-api"
	"gitlab.login.no/tekkom/web/beehive/admin-api/docs"
)

func (server *Server) setSwaggerInfo() {
	docs.SwaggerInfo.Title = "Beehive Admin API"
	docs.SwaggerInfo.Description = "Admin API for Beehive"
	docs.SwaggerInfo.Version = adminapi.Version()
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
