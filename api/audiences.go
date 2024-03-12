package api

import (
	"net/http"

	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
	"github.com/gin-gonic/gin"
)

// getAudiences godoc
//
//	@ID				get-audiences
//	@Summary		Get all audiences
//	@Description	Get all audiences
//	@Tags			audiences
//	@Produce		json
//	@Success		200	{array}		db.GetAudiencesRow
//	@Failure		500	{object}	errorResponse
//	@Router			/audiences [get]
func (server *Server) getAudiences(ctx *gin.Context) {
	audiences, err := server.service.GetAudiences(ctx)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, audiences)
}

type getAudienceRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// getAudience godoc
//
//	@ID				get-audience
//	@Summary		Get audience by ID
//	@Description	Get audience by ID
//	@Tags			audiences
//	@Produce		json
//	@Param			params	path		getAudienceRequest	true	"Audience ID"
//	@Success		200		{object}	db.Audience
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/audiences/{id} [get]
func (server *Server) getAudience(ctx *gin.Context) {
	var req getAudienceRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	audience, err := server.service.GetAudience(ctx, req.ID)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, audience)
}
