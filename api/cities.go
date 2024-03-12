package api

import (
	"net/http"

	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
	"github.com/gin-gonic/gin"
)

// getAllCities godoc
//
//	@ID				get-all-cities
//	@Summary		Get all cities
//	@Description	Get all cities
//	@Tags			cities
//	@Produce		json
//	@Success		200	{array}		string
//	@Failure		500	{object}	errorResponse
//	@Router			/cities [get]
func (server *Server) getAllCities(ctx *gin.Context) {
	cities, err := server.service.GetAllCities(ctx)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, cities)
}
