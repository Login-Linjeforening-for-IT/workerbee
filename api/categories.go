package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getCategories(ctx *gin.Context) {
	categories, err := server.service.GetCategories(ctx)
	if err != nil {
		server.writeError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

type getCategoryRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getCategory(ctx *gin.Context) {
	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	category, err := server.service.GetCategory(ctx, req.ID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, category)
}
