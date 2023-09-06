package api

import (
	"database/sql"
	"net/http"
	"time"

	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null/zero"
)

type getJobsRequest struct {
	Limit  int32 `form:"limit,default=20"`
	Offset int32 `form:"offset"`
}

func (server *Server) getJobs(ctx *gin.Context) {
	var req getJobsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	jobs, err := server.service.GetJobs(ctx, db.GetJobsParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		server.writeError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, jobs)
}

type getJobRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getJob(ctx *gin.Context) {
	var req getJobRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	job, err := server.service.GetJob(ctx, req.ID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, job)
}

type createJobRequest struct {
	Visible             bool        `json:"visible"`
	TitleNo             string      `json:"title_no" binding:"required"`
	TitleEn             zero.String `json:"title_en"`
	PositionTitleNo     string      `json:"position_title_no" binding:"required"`
	PositionTitleEn     zero.String `json:"position_title_en"`
	DescriptionShortNo  string      `json:"description_short_no" binding:"required"`
	DescriptionShortEn  zero.String `json:"description_short_en"`
	DescriptionLongNo   string      `json:"description_long_no" binding:"required"`
	DescriptionLongEn   zero.String `json:"description_long_en"`
	JobType             db.JobType  `json:"job_type" binding:"required,jobtype"`
	TimePublish         time.Time   `json:"time_publish" binding:"required"`
	ApplicationDeadline time.Time   `json:"application_deadline" binding:"required"`
	BannerImage         zero.String `json:"banner_image"`
	Organization        string      `json:"organization"`
	ApplicationURL      zero.String `json:"application_url"` // TODO: Make nullable in db

	// TODO: consider adding skills and audiences here
}

func (server *Server) createJob(ctx *gin.Context) {
	var req createJobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[createJobRequest](server, ctx, err)
		return
	}

	job, err := server.service.CreateJob(ctx, db.CreateJobParams{
		Visible:             req.Visible,
		TitleNo:             req.TitleNo,
		TitleEn:             req.TitleEn,
		PositionTitleNo:     req.PositionTitleNo,
		PositionTitleEn:     req.PositionTitleEn,
		DescriptionShortNo:  req.DescriptionShortNo,
		DescriptionShortEn:  req.DescriptionShortEn,
		DescriptionLongNo:   req.DescriptionLongNo,
		DescriptionLongEn:   req.DescriptionLongEn,
		JobType:             req.JobType,
		TimePublish:         req.TimePublish,
		ApplicationDeadline: req.ApplicationDeadline,
		BannerImage:         req.BannerImage,
		Organization:        req.Organization,
		// ApplicationURL:      req.ApplicationURL.String,
	})
	if err != nil {
		server.writeError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, job)
}

type updateJobRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
	db.UpdateJobParams
}

func (server *Server) updateJob(ctx *gin.Context) {
	var req updateJobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[updateJobRequest](server, ctx, err)
		return
	}

	job, err := server.service.UpdateJob(ctx, req.UpdateJobParams)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, job)
}

type deleteJobRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteJob(ctx *gin.Context) {
	var req deleteJobRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	job, err := server.service.SoftDeleteJob(ctx, req.ID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, job)
}
