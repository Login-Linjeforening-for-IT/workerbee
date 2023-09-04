package api

import (
	"database/sql"
	"net/http"

	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
	"github.com/gin-gonic/gin"
)

type getEventsRequest struct {
	Limit      int32 `form:"limit"`
	Offset     int32 `form:"offset"`
	Historical bool  `form:"historical"`
}

func (server *Server) getEvents(ctx *gin.Context) {
	var req getEventsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	events, err := server.service.GetEvents(ctx, db.GetEventsParams{
		Historical: req.Historical,
		Offset:     req.Offset,
		Limit:      req.Limit,
	})
	if err != nil {
		server.writeError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, events)
}

type getEventRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getEvent(ctx *gin.Context) {
	var req getEventRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	event, err := server.service.GetEventDetails(ctx, req.ID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, event)
}
