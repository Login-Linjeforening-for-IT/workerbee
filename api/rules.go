package api

import (
	"database/sql"
	"net/http"

	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null/zero"
)

type getRulesRequest struct {
	Limit  int32 `form:"limit,default=20"`
	Offset int32 `form:"offset"`
}

func (server *Server) getRules(ctx *gin.Context) {
	var req getRulesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	rules, err := server.service.GetRules(ctx, db.GetRulesParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		server.writeError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, rules)
}

type getRuleRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getRule(ctx *gin.Context) {
	var req getRuleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	rule, err := server.service.GetRule(ctx, req.ID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, rule)
}

type createRuleRequest struct {
	NameNo        string      `json:"name_no" binding:"required"`
	NameEn        zero.String `json:"name_en"`
	DescriptionNo string      `json:"description_no" binding:"required"`
	DescriptionEn zero.String `json:"description_en"`
}

func (server *Server) createRule(ctx *gin.Context) {
	var req createRuleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// server.writeError(ctx, http.StatusBadRequest, err)
		writeValidationError[createEventRequest](server, ctx, err)
		return
	}

	rule, err := server.service.CreateRule(ctx, db.CreateRuleParams{
		NameNo:        req.NameNo,
		NameEn:        req.NameEn,
		DescriptionNo: req.DescriptionNo,
		DescriptionEn: req.DescriptionEn,
	})
	if err != nil {
		server.writeError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, rule)
}

type updateRuleRequest struct {
	ID int32 `json:"id" binding:"required,min=1"`
	db.UpdateRuleParams
}

func (server *Server) updateRule(ctx *gin.Context) {
	var req updateRuleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[updateRuleRequest](server, ctx, err)
		return
	}

	req.UpdateRuleParams.ID = req.ID

	rule, err := server.service.UpdateRule(ctx, req.UpdateRuleParams)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, rule)
}
