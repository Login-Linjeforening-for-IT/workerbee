package api

import (
	"database/sql"
	"net/http"
	"time"

	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null/zero"
)

type getEventsRequest struct {
	Limit      int32 `form:"limit,default=20"`
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

type createEventRequest struct {
	Visible bool `json:"visible"`

	NameNo string      `json:"name_no" binding:"required"`
	NameEn zero.String `json:"name_en"`

	DescriptionNo string      `json:"description_no" binding:"required"`
	DescriptionEn zero.String `json:"description_en"`

	InformationalNo zero.String `json:"informational_no"`
	InformationalEn zero.String `json:"informational_en"`

	TimeStart          time.Time `json:"time_start" binding:"required"`
	TimeEnd            zero.Time `json:"time_end"`
	TimePublish        zero.Time `json:"time_publish"`
	TimeSignupRelease  zero.Time `json:"time_signup_release"`
	TimeSignupDeadline zero.Time `json:"time_signup_deadline"`

	Canceled  bool `json:"canceled"`
	Digital   bool `json:"digital"`
	Highlight bool `json:"highlight"`

	ImageSmall  string `json:"image_small"`
	ImageBanner string `json:"image_banner"`

	LinkFacebook string      `json:"link_facebook" binding:"required"`
	LinkDiscord  string      `json:"link_discord" binding:"required"` // TODO: should this be optional?
	LinkSignup   string      `json:"link_signup" binding:"required"`  // TODO: should this be optional?
	LinkStream   zero.String `json:"link_stream"`

	Capacity zero.Int `json:"capacity"`
	Full     bool     `json:"full"`

	Category int32    `json:"category" binding:"required"`
	Location zero.Int `json:"location"`
	Parent   zero.Int `json:"parent"`
	Rule     zero.Int `json:"rule"`

	/* // TODO: should these be here or should they be as a patch?
	Organizations []string `json:"organizations"`
	Audiences     []int    `json:"audiences"`
	*/
}

func (server *Server) createEvent(ctx *gin.Context) {
	var req createEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[createEventRequest](server, ctx, err)
		return
	}

	// TODO: time type stuff
	// Is time type required?
	// time type is inferred

	event, err := server.service.CreateEvent(ctx, db.CreateEventParams{
		Visible:            req.Visible,
		NameNo:             req.NameNo,
		NameEn:             req.NameEn,
		DescriptionNo:      req.DescriptionNo,
		DescriptionEn:      req.DescriptionEn,
		InformationalNo:    req.InformationalNo,
		InformationalEn:    req.InformationalEn,
		TimeType:           "", // TODO
		TimeStart:          req.TimeStart,
		TimeEnd:            req.TimeEnd,
		TimePublish:        req.TimePublish,
		TimeSignupRelease:  req.TimeSignupRelease,
		TimeSignupDeadline: req.TimeSignupDeadline,
		Canceled:           req.Canceled,
		Digital:            req.Digital,
		Highlight:          req.Highlight,
		ImageSmall:         req.ImageSmall,
		ImageBanner:        req.ImageBanner,
		LinkFacebook:       req.LinkFacebook,
		LinkDiscord:        req.LinkDiscord,
		LinkSignup:         req.LinkSignup,
		LinkStream:         req.LinkStream,
		Capacity:           req.Capacity,
		Full:               req.Full,
		Category:           req.Category,
		Location:           req.Location,
		Parent:             req.Parent,
		Rule:               req.Rule,
	})
	if err != nil {
		// TODO: check duplicate (should not be possible), bad values, etc.
		server.writeError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, event)
}

type updateEventRequest struct {
	ID int32 `json:"id" binding:"required"`
	db.UpdateEventParams
}

func (server *Server) updateEvent(ctx *gin.Context) {
	var req updateEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[updateEventRequest](server, ctx, err)
		return
	}

	// TODO: Check enum validity

	req.UpdateEventParams.ID = req.ID

	event, err := server.service.UpdateEvent(ctx, req.UpdateEventParams)
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
