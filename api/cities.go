package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getAllCities(ctx *gin.Context) {
	cities, err := server.service.GetAllCities(ctx)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, cities)
}
