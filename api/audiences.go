package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getAudiences(ctx *gin.Context) {
	audiences, err := server.service.GetAudiences(ctx)
	if err != nil {
		server.writeError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, audiences)
}

type getAudienceRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAudience(ctx *gin.Context) {
	var req getAudienceRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	audience, err := server.service.GetAudience(ctx, req.ID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, audience)
}
