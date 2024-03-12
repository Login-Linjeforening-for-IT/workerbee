package api

import (
	"net/http"

	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null/zero"
)

type getOrganizationsRequest struct {
	Limit  int32 `form:"limit,default=20"`
	Offset int32 `form:"offset"`
}

// getOrganizations godoc
//
//	@Summary		Get organizations
//	@Description	Get a list of organizations
//	@Tags			organizations
//	@Produce		json
//	@Param			params	query		getOrganizationsRequest	false	"Parameters"
//	@Success		200		{array}		db.GetOrganizationsRow	"OK"
//	@Failure		400		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/organizations [get]
func (server *Server) getOrganizations(ctx *gin.Context) {
	var req getOrganizationsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	organizations, err := server.service.GetOrganizations(ctx, db.GetOrganizationsParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, organizations)
}

type getOrganizationRequest struct {
	Shortname string `uri:"shortname" binding:"required,min=1"`
}

// getOrganization godoc
//
//	@Summary		Get organization
//	@Description	Get an organization by shortname
//	@Tags			organizations
//	@Produce		json
//	@Param			shortname	path		string	true	"Shortname"
//	@Success		200			{object}	db.Organization
//	@Failure		400			{object}	errorResponse
//	@Failure		404			{object}	errorResponse
//	@Failure		500			{object}	errorResponse
//	@Router			/organizations/{shortname} [get]
func (server *Server) getOrganization(ctx *gin.Context) {
	var req getOrganizationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	organization, err := server.service.GetOrganization(ctx, req.Shortname)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, organization)
}

type createOrganizationRequest struct {
	Shortname     string      `json:"shortname" binding:"required,min=1"`
	NameNo        string      `json:"name_no" binding:"required,min=1"`
	NameEn        zero.String `json:"name_en"`
	DescriptionNo string      `json:"description_no"`
	DescriptionEn zero.String `json:"description_en"`
	Type          zero.Int    `json:"type"`
	LinkHomepage  zero.String `json:"link_homepage"`
	LinkLinkedin  zero.String `json:"link_linkedin"`
	LinkFacebook  zero.String `json:"link_facebook"`
	LinkInstagram zero.String `json:"link_instagram"`
	Logo          zero.String `json:"logo"`
}

// createOrganization godoc
//
//	@Summary		Create organization
//	@Description	Create a new organization
//	@Tags			organizations
//	@Accept			json
//	@Produce		json
//	@Param			request	body		createOrganizationRequest	true	"Request"
//	@Success		200		{object}	db.Organization
//	@Failure		400		{object}	errorResponse
//	@Failure		409		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/organizations [post]
func (server *Server) createOrganization(ctx *gin.Context) {
	var req createOrganizationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[createOrganizationRequest](server, ctx, err)
		return
	}

	organization, err := server.service.CreateOrganization(ctx, db.CreateOrganizationParams{
		Shortname:     req.Shortname,
		NameNo:        req.NameNo,
		NameEn:        req.NameEn,
		DescriptionNo: req.DescriptionNo,
		DescriptionEn: req.DescriptionEn,
		Type:          int32(req.Type.Int64),
		LinkHomepage:  req.LinkHomepage,
		LinkLinkedin:  req.LinkLinkedin,
		LinkFacebook:  req.LinkFacebook,
		LinkInstagram: req.LinkInstagram,
		Logo:          req.Logo,
	})
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, organization)
}

type updateOrganizationRequest struct {
	Shortname string `json:"shortname" binding:"required,min=1"`
	db.UpdateOrganizationParams
}

// updateOrganization godoc
//
//	@Summary		Update organization
//	@Description	Update an organization by shortname
//	@Tags			organizations
//	@Accept			json
//	@Produce		json
//	@Param			shortname	path		string						true	"Shortname"
//	@Param			request		body		db.UpdateOrganizationParams	true	"Request"
//	@Success		200			{object}	db.Organization
//	@Failure		400			{object}	errorResponse
//	@Failure		500			{object}	errorResponse
//	@Router			/organizations/{shortname} [patch]
func (server *Server) updateOrganization(ctx *gin.Context) {
	var req updateOrganizationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[updateOrganizationRequest](server, ctx, err)
		return
	}

	req.UpdateOrganizationParams.Shortname = req.Shortname

	organization, err := server.service.UpdateOrganization(ctx, req.UpdateOrganizationParams)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, organization)
}

type deleteOrganizationRequest struct {
	Shortname string `uri:"shortname" binding:"required,min=1"`
}

// deleteOrganization godoc
//
//	@Summary		Delete organization
//	@Description	Delete an organization by shortname
//	@Tags			organizations
//	@Produce		json
//	@Param			shortname	path		string	true	"Shortname"
//	@Success		200			{object}	db.Organization
//	@Failure		400			{object}	errorResponse
//	@Failure		500			{object}	errorResponse
//	@Router			/organizations/{shortname} [delete]
func (server *Server) deleteOrganization(ctx *gin.Context) {
	var req deleteOrganizationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	organization, err := server.service.SoftDeleteOrganization(ctx, req.Shortname)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, organization)
}
