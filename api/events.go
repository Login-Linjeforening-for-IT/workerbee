package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null/zero"
	db "gitlab.login.no/tekkom/web/beehive/admin-api/db/sqlc"
)

type getEventsRequest struct {
	Limit      int32 `form:"limit,default=20"`
	Offset     int32 `form:"offset"`
	Historical bool  `form:"historical"`
}

// getEvents godoc
//
//	@ID				get-events
//	@Summary		Get all events
//	@Description	Get all events
//	@Tags			events
//	@Produce		json
//	@Param			params	query		getEventsRequest	true	"Event parameters"
//	@Success		200		{array}		db.GetEventsRow
//	@Failure		400		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/events [get]
func (server *Server) getEvents(ctx *gin.Context) {
	var req getEventsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, fmt.Errorf("getEvents, ShouldBindQuery - %w", err))
		return
	}

	events, err := server.service.GetEvents(ctx, db.GetEventsParams{
		Historical: req.Historical,
		Offset:     req.Offset,
		Limit:      req.Limit,
	})
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, events)
}

type getEventRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// getEvent godoc
//
//	@ID				get-event
//	@Summary		Get event by ID
//	@Description	Get event by ID
//	@Tags			events
//	@Produce		json
//	@Param			params	path		getEventRequest	true	"Event ID"
//	@Success		200		{object}	service.EventDetails
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/events/{id} [get]
func (server *Server) getEvent(ctx *gin.Context) {
	var req getEventRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, fmt.Errorf("getEvent, ShouldBindUri - %w", err))
		return
	}

	event, err := server.service.GetEventDetails(ctx, req.ID)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
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

// createEvent godoc
//
//	@ID				create-event
//	@Summary		Create event
//	@Description	Create event
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			params	body		createEventRequest	true	"Event parameters"
//	@Success		201		{object}	db.Event
//	@Failure		400		{object}	errorResponse	"Invalid input"
//	@Failure		404		{object}	errorResponse	"Foreign key violation"
//	@Failure		409		{object}	errorResponse	"Event already exists"
//	@Failure		500		{object}	errorResponse
//	@Router			/events [post]
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
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, event)
}

type updateEventRequest struct {
	ID int32 `json:"id" binding:"required"`
	db.UpdateEventParams
}

// updateEvent godoc
//
//	@ID				update-event
//	@Summary		Update event
//	@Description	Update event
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			params	body		updateEventRequest	true	"Event parameters"
//	@Success		200		{object}	db.Event
//	@Failure		400		{object}	errorResponse	"Invalid input"
//	@Failure		404		{object}	errorResponse	"Foreign key violation"
//	@Failure		500		{object}	errorResponse
//	@Router			/events [patch]
func (server *Server) updateEvent(ctx *gin.Context) {
	var req updateEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[updateEventRequest](server, ctx, err)
		return
	}

	if req.TimeType.Valid && !(req.TimeType.TimeTypeEnum).Valid() {
		server.writeError(ctx, http.StatusBadRequest, fmt.Errorf("updateEvent, ShouldBindJSON - %w", errors.New("invalid time type")))
		return
	}

	req.UpdateEventParams.ID = req.ID

	event, err := server.service.UpdateEvent(ctx, req.UpdateEventParams)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, event)
}

type deleteEventRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// deleteEvent godoc
//
//	@ID				delete-event
//	@Summary		Delete event by ID
//	@Description	Delete event by ID
//	@Tags			events
//	@Produce		json
//	@Param			params	path		deleteEventRequest	true	"Event ID"
//	@Success		200		{object}	db.Event
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/events/{id} [delete]
func (server *Server) deleteEvent(ctx *gin.Context) {
	var req deleteEventRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, fmt.Errorf("deleteEvent, ShouldBindUri - %w", err))
		return
	}

	event, err := server.service.SoftDeleteEvent(ctx, req.ID)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, event)
}

type addOrganizationToEventRequest struct {
	Event        int32  `json:"event" binding:"required,min=1"`
	Organization string `json:"organization" binding:"required"`
}

// addOrganizationToEvent godoc
//
//	@ID				add-organization-to-event
//	@Summary		Add organization to event
//	@Description	Add organization to event
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			params	body		addOrganizationToEventRequest	true	"Organization parameters"
//	@Success		204		{object}	nil
//	@Failure		400		{object}	errorResponse	"Invalid input"
//	@Failure		404		{object}	errorResponse	"Foreign key violation"
//	@Failure		409		{object}	errorResponse	"Organization already exists"
//	@Failure		500		{object}	errorResponse
//	@Router			/events/organizations [post]
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
		server.writeDBError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

type removeOrganizationFromEventRequest struct {
	Event        int32  `json:"event" binding:"required,min=1"`
	Organization string `json:"organization" binding:"required"`
}

// removeOrganizationFromEvent godoc
//
//	@ID				remove-organization-from-event
//	@Summary		Remove organization from event
//	@Description	Remove organization from event
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			params	body		removeOrganizationFromEventRequest	true	"Organization parameters"
//	@Success		204		{object}	nil
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/events/organizations [delete]
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
		server.writeDBError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

type addAudienceToEventRequest struct {
	Event    int32 `json:"event" binding:"required,min=1"`
	Audience int32 `json:"audience" binding:"required,min=1"`
}

// addAudienceToEvent godoc
//
//	@ID				add-audience-to-event
//	@Summary		Add audience to event
//	@Description	Add audience to event
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			params	body		addAudienceToEventRequest	true	"Audience parameters"
//	@Success		204		{object}	nil
//	@Failure		400		{object}	errorResponse	"Invalid input"
//	@Failure		404		{object}	errorResponse	"Foreign key violation"
//	@Failure		409		{object}	errorResponse	"Audience already exists"
//	@Failure		500		{object}	errorResponse
//	@Router			/events/audiences [post]
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
		server.writeDBError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

type removeAudienceFromEventRequest struct {
	Event    int32 `json:"event" binding:"required,min=1"`
	Audience int32 `json:"audience" binding:"required,min=1"`
}

// removeAudienceFromEvent godoc
//
//	@ID				remove-audience-from-event
//	@Summary		Remove audience from event
//	@Description	Remove audience from event
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			params	body		removeAudienceFromEventRequest	true	"Audience parameters"
//	@Success		204		{object}	nil
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/events/audiences [delete]
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
		server.writeDBError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
