package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null/zero"
	db "gitlab.login.no/tekkom/web/beehive/admin-api/db/sqlc"
)

type getJobsRequest struct {
	Limit  int32 `form:"limit,default=20"`
	Offset int32 `form:"offset"`
}

// getJobs godoc
//
//	@Summary		Get jobs
//	@Description	Get a list of jobs
//	@Tags			jobs
//	@Produce		json
//	@Param			limit	query		int	false	"Limit"
//	@Param			offset	query		int	false	"Offset"
//	@Success		200		{array}		db.GetJobsRow
//	@Failure		400		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/jobs [get]
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
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, jobs)
}

type getJobRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// getJob godoc
//
//	@Summary		Get job
//	@Description	Get a job by ID
//	@Tags			jobs
//	@Produce		json
//	@Param			id	path		int	true	"Job ID"
//	@Success		200	{object}	db.GetJobRow
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/jobs/{id} [get]
func (server *Server) getJob(ctx *gin.Context) {
	var req getJobRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	job, err := server.service.GetJob(ctx, req.ID)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, job)
}

type createJobRequest struct {
	Visible             bool        `json:"visible"`
	Highlight           bool        `json:"highlight"`
	TitleNo             string      `json:"title_no" binding:"required"`
	TitleEn             zero.String `json:"title_en"`
	PositionTitleNo     string      `json:"position_title_no" binding:"required"`
	PositionTitleEn     zero.String `json:"position_title_en"`
	DescriptionShortNo  string      `json:"description_short_no" binding:"required"`
	DescriptionShortEn  zero.String `json:"description_short_en"`
	DescriptionLongNo   string      `json:"description_long_no" binding:"required"`
	DescriptionLongEn   zero.String `json:"description_long_en"`
	JobType             db.JobType  `json:"job_type" binding:"required"`
	TimePublish         time.Time   `json:"time_publish" binding:"required"`
	TimeExpire          time.Time   `json:"time_expire" binding:"required"`
	ApplicationDeadline time.Time   `json:"application_deadline" binding:"required"`
	BannerImage         zero.String `json:"banner_image"`
	Organization        string      `json:"organization"`
	ApplicationURL      zero.String `json:"application_url"`

	// TODO: consider adding skills and audiences here
}

// createJob godoc
//
//	@Summary		Create job
//	@Description	Create a new job
//	@Tags			jobs
//	@Produce		json
//	@Param			request	body		createJobRequest	true	"Job details"
//	@Success		200		{object}	db.GetJobRow
//	@Failure		400		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/jobs [post]
func (server *Server) createJob(ctx *gin.Context) {
	var req createJobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[createJobRequest](server, ctx, err)
		return
	}

	job, err := server.service.CreateJob(ctx, db.CreateJobParams{
		Visible:             req.Visible,
		Highlight:           req.Highlight,
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
		TimeExpire:          req.TimeExpire,
		ApplicationDeadline: req.ApplicationDeadline,
		BannerImage:         req.BannerImage,
		Organization:        req.Organization,
		ApplicationUrl:      req.ApplicationURL,
	})
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, job)
}

type updateJobRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
	db.UpdateJobParams
}

// updateJob godoc
//
//	@Summary		Update job
//	@Description	Update a job by ID
//	@Tags			jobs
//	@Produce		json
//	@Param			id		path		int					true	"Job ID"
//	@Param			request	body		db.UpdateJobParams	true	"Job details"
//	@Success		200		{object}	db.GetJobRow
//	@Failure		400		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/jobs [patch]
func (server *Server) updateJob(ctx *gin.Context) {
	var req updateJobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[updateJobRequest](server, ctx, err)
		return
	}

	req.UpdateJobParams.ID = req.ID

	job, err := server.service.UpdateJob(ctx, req.UpdateJobParams)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, job)
}

type deleteJobRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// deleteJob godoc
//
//	@Summary		Delete job
//	@Description	Delete a job by ID
//	@Tags			jobs
//	@Produce		json
//	@Param			id	path		int	true	"Job ID"
//	@Success		200	{object}	db.GetJobRow
//	@Failure		400	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/jobs/{id} [delete]
func (server *Server) deleteJob(ctx *gin.Context) {
	var req deleteJobRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	job, err := server.service.SoftDeleteJob(ctx, req.ID)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, job)
}

type addSkillToJobRequest struct {
	JobID int32  `json:"id" binding:"required,min=1"`
	Skill string `json:"skill" binding:"required,min=1"`
}

// addSkillToJob godoc
//
//	@Summary		Add skill to job
//	@Description	Add a skill to a job
//	@Tags			jobs
//	@Produce		json
//	@Param			body	body		addSkillToJobRequest	true	"Job ID and skill"
//	@Success		204		{object}	nil
//	@Failure		400		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/jobs/skills [post]
func (server *Server) addSkillToJob(ctx *gin.Context) {
	var req addSkillToJobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[addSkillToJobRequest](server, ctx, err)
		return
	}

	err := server.service.AddSkillToJob(ctx, db.AddSkillToJobParams{
		Ad:    req.JobID,
		Skill: req.Skill,
	})
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

type removeSkillFromJobRequest struct {
	JobID int32  `json:"id" binding:"required,min=1"`
	Skill string `json:"skill" binding:"required,min=1"`
}

// removeSkillFromJob godoc
//
//	@Summary		Remove skill from job
//	@Description	Remove a skill from a job
//	@Tags			jobs
//	@Produce		json
//	@Param			body	body		removeSkillFromJobRequest	true	"Job ID and skill"
//	@Success		204		{object}	nil
//	@Failure		400		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
func (server *Server) removeSkillFromJob(ctx *gin.Context) {
	var req removeSkillFromJobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[removeSkillFromJobRequest](server, ctx, err)
		return
	}

	err := server.service.RemoveSkillFromJob(ctx, db.RemoveSkillFromJobParams{
		Ad:    req.JobID,
		Skill: req.Skill,
	})
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

type addCityToJobRequest struct {
	JobID int32  `json:"id" binding:"required,min=1"`
	City  string `json:"city" binding:"required,min=1"`
}

// addCityToJob godoc
//
//	@Summary		Add city to job
//	@Description	Add a city to a job
//	@Tags			jobs
//	@Produce		json
//	@Param			body	body		addCityToJobRequest	true	"Job ID and city"
//	@Success		204		{object}	nil
//	@Failure		400		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/jobs/cities [post]
func (server *Server) addCityToJob(ctx *gin.Context) {
	var req addCityToJobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[addCityToJobRequest](server, ctx, err)
		return
	}

	err := server.service.AddCityToJob(ctx, db.AddCityToJobParams{
		Ad:   req.JobID,
		City: req.City,
	})
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

type removeCityFromJobRequest struct {
	JobID int32  `json:"id" binding:"required,min=1"`
	City  string `json:"city" binding:"required,min=1"`
}

// removeCityFromJob godoc
//
//	@Summary		Remove city from job
//	@Description	Remove a city from a job
//	@Tags			jobs
//	@Produce		json
//	@Param			body	body		removeCityFromJobRequest	true	"Job ID and city"
//	@Success		204		{object}	nil
//	@Failure		400		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/jobs/cities [delete]
func (server *Server) removeCityFromJob(ctx *gin.Context) {
	var req removeCityFromJobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[removeCityFromJobRequest](server, ctx, err)
		return
	}

	err := server.service.RemoveCityFromJob(ctx, db.RemoveCityFromJobParams{
		Ad:   req.JobID,
		City: req.City,
	})
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
