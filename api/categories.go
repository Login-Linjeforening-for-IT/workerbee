package api

import (
	"net/http"

	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
	"github.com/gin-gonic/gin"
)

// getCategories godoc
//
//	@ID				get-categories
//	@Summary		Get all categories
//	@Description	Get all categories
//	@Tags			categories
//	@Produce		json
//	@Success		200	{array}		db.GetCategoriesRow
//	@Failure		500	{object}	errorResponse
//	@Router			/categories [get]
func (server *Server) getCategories(ctx *gin.Context) {
	categories, err := server.service.GetCategories(ctx)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

type getCategoryRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// getCategory godoc
//
//	@ID				get-category
//	@Summary		Get category by ID
//	@Description	Get category by ID
//	@Tags			categories
//	@Produce		json
//	@Param			params	path		getCategoryRequest	true	"Category ID"
//	@Success		200		{object}	db.Category
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
func (server *Server) getCategory(ctx *gin.Context) {
	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	category, err := server.service.GetCategory(ctx, req.ID)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, category)
}
