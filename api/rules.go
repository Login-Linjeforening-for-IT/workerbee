package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null/zero"
	db "gitlab.login.no/tekkom/web/beehive/admin-api/db/sqlc"
)

type getRulesRequest struct {
	Limit  int32 `form:"limit,default=20"`
	Offset int32 `form:"offset"`
}

// getRules godoc
//
//	@Summary		Get rules
//	@Description	Get a list of rules
//	@Tags			rules
//	@Produce		json
//	@Param			params	query		getRulesRequest	false	"Parameters"
//	@Success		200		{array}		db.GetRulesRow	"OK"
//	@Failure		400		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/rules [get]
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
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, rules)
}

type getRuleRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// getRule godoc
//
//	@Summary		Get rule
//	@Description	Get a rule by ID
//	@Tags			rules
//	@Produce		json
//	@Param			id	path		int	true	"Rule ID"
//	@Success		200	{object}	db.Rule
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/rules/{id} [get]
func (server *Server) getRule(ctx *gin.Context) {
	var req getRuleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	rule, err := server.service.GetRule(ctx, req.ID)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
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

// createRule godoc
//
//	@Summary		Create rule
//	@Description	Create a new rule
//	@Tags			rules
//	@Accept			json
//	@Produce		json
//	@Param			params	body		createRuleRequest	true	"Parameters"
//	@Success		200		{object}	db.Rule
//	@Failure		400		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/rules [post]
func (server *Server) createRule(ctx *gin.Context) {
	var req createRuleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[createEventRequest](server, ctx, err)
		return
	}

	rule, err := server.service.CreateRule(ctx, db.CreateRuleParams{
		NameNo:        req.NameNo,
		NameEn:        req.NameEn,
		DescriptionNo: req.DescriptionNo,
		DescriptionEn: req.DescriptionEn,
	})
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
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
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, rule)
}

type deleteRuleRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteRule(ctx *gin.Context) {
	var req deleteRuleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	rule, err := server.service.SoftDeleteRule(ctx, req.ID)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, rule)
}
