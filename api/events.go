package api

import (
	"database/sql"
	"errors"
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

	TimeType           db.TimeTypeEnum `json:"time_type" binding:"required,timetypeenum"`
	TimeStart          time.Time       `json:"time_start" binding:"required"`
	TimeEnd            time.Time       `json:"time_end"`
	TimePublish        zero.Time       `json:"time_publish"`
	TimeSignupRelease  zero.Time       `json:"time_signup_release"`
	TimeSignupDeadline zero.Time       `json:"time_signup_deadline"`

	Canceled  bool `json:"canceled"`
	Digital   bool `json:"digital"`
	Highlight bool `json:"highlight"`

	ImageSmall  zero.String `json:"image_small"`
	ImageBanner string      `json:"image_banner"`

	LinkFacebook zero.String `json:"link_facebook" binding:"required"`
	LinkDiscord  zero.String `json:"link_discord" binding:"required"`
	LinkSignup   zero.String `json:"link_signup" binding:"required"`
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

	event, err := server.service.CreateEvent(ctx, db.CreateEventParams{
		Visible:            req.Visible,
		NameNo:             req.NameNo,
		NameEn:             req.NameEn,
		DescriptionNo:      req.DescriptionNo,
		DescriptionEn:      req.DescriptionEn,
		InformationalNo:    req.InformationalNo,
		InformationalEn:    req.InformationalEn,
		TimeType:           req.TimeType,
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
	err = db.ParseError(err)
	if err != nil {
		switch err.(type) {
		case *db.ForeignKeyViolationError, *db.NotFoundError:
			server.writeError(ctx, http.StatusNotFound, err)
		case *db.UniqueViolationError:
			server.writeError(ctx, http.StatusConflict, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
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

	if req.TimeType.Valid && !db.IsValidTimeTypeEnum(req.TimeType.TimeTypeEnum) {
		server.writeError(ctx, http.StatusBadRequest, errors.New("invalid time type"))
		return
	}

	req.UpdateEventParams.ID = req.ID

	event, err := server.service.UpdateEvent(ctx, req.UpdateEventParams)
	err = db.ParseError(err)
	if err != nil {
		switch err.(type) {
		case *db.ForeignKeyViolationError, *db.NotFoundError:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)

		}
		return
	}

	ctx.JSON(http.StatusOK, event)
}

type deleteEventRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteEvent(ctx *gin.Context) {
	var req deleteEventRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	event, err := server.service.SoftDeleteEvent(ctx, req.ID)
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

type addOrganizationToEventRequest struct {
	Event        int32  `json:"event" binding:"required,min=1"`
	Organization string `json:"organization" binding:"required"`
}

func (server *Server) addOrganizationToEvent(ctx *gin.Context) {
	var req addOrganizationToEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[addOrganizationToEventRequest](server, ctx, err)
		return
	}

	err := server.service.AddOrganizationToEvent(ctx, db.AddOrganizationToEventParams{
		Event:        req.Event,
		Organization: req.Organization,
	})
	err = db.ParseError(err)
	if err != nil {
		switch err.(type) {
		case *db.ForeignKeyViolationError, *db.NotFoundError:
			server.writeError(ctx, http.StatusNotFound, err)
		case *db.UniqueViolationError:
			server.writeError(ctx, http.StatusConflict, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.Status(http.StatusCreated)
}

type removeOrganizationFromEventRequest struct {
	Event        int32  `json:"event" binding:"required,min=1"`
	Organization string `json:"organization" binding:"required"`
}

func (server *Server) removeOrganizationFromEvent(ctx *gin.Context) {
	var req removeOrganizationFromEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[removeOrganizationFromEventRequest](server, ctx, err)
		return
	}

	err := server.service.RemoveOrganizationFromEvent(ctx, db.RemoveOrganizationFromEventParams{
		Event:        req.Event,
		Organization: req.Organization,
	})
	err = db.ParseError(err)
	if err != nil {
		switch err.(type) {
		case *db.ForeignKeyViolationError, *db.NotFoundError:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.Status(http.StatusOK)
}

type addAudienceToEventRequest struct {
	Event    int32 `json:"event" binding:"required,min=1"`
	Audience int32 `json:"audience" binding:"required,min=1"`
}

func (server *Server) addAudienceToEvent(ctx *gin.Context) {
	var req addAudienceToEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[addAudienceToEventRequest](server, ctx, err)
		return
	}

	err := server.service.AddAudienceToEvent(ctx, db.AddAudienceToEventParams{
		Event:    req.Event,
		Audience: req.Audience,
	})
	err = db.ParseError(err)
	if err != nil {
		switch err.(type) {
		case *db.ForeignKeyViolationError, *db.NotFoundError:
			server.writeError(ctx, http.StatusNotFound, err)
		case *db.UniqueViolationError:
			server.writeError(ctx, http.StatusConflict, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.Status(http.StatusCreated)
}

type removeAudienceFromEventRequest struct {
	Event    int32 `json:"event" binding:"required,min=1"`
	Audience int32 `json:"audience" binding:"required,min=1"`
}

func (server *Server) removeAudienceFromEvent(ctx *gin.Context) {
	var req removeAudienceFromEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[removeAudienceFromEventRequest](server, ctx, err)
		return
	}

	err := server.service.RemoveAudienceFromEvent(ctx, db.RemoveAudienceFromEventParams{
		Event:    req.Event,
		Audience: req.Audience,
	})
	err = db.ParseError(err)
	if err != nil {
		switch err.(type) {
		case *db.ForeignKeyViolationError, *db.NotFoundError:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.Status(http.StatusOK)
}
