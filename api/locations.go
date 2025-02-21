package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null/zero"
	db "gitlab.login.no/tekkom/web/beehive/admin-api/db/sqlc"
)

type getLocationsRequest struct {
	Limit  int32  `form:"limit,default=200"`
	Offset int32  `form:"offset"`
	Type   string `form:"type"`
}

// getLocations godoc
//
//	@ID				get-locations
//	@Summary		Get locations
//	@Description	Get a list of locations
//	@Tags			locations
//	@Produce		json
//	@Param			params	query		getLocationsRequest			false	"Parameters"
//	@Success		200		{array}		db.GetMazemapLocationsRow	"OK - mazemap"
//	@Success		200		{array}		db.GetCoordsLocationsRow	"OK - coords"
//	@Success		200		{array}		db.GetAddressLocationsRow	"OK - address"
//	@Success		200		{array}		db.GetLocationsRow			"OK - other"
//	@Failure		400		{object}	errorResponse
//	@Failure		404		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/locations [get]
func (server *Server) getLocations(ctx *gin.Context) {
	var req getLocationsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, fmt.Errorf("getLocations, ShouldBindQuery - %w", err))
		return
	}

	var locations any
	var err error

	switch req.Type {
	case "mazemap":
		locations, err = server.service.GetMazemapLocations(ctx, db.GetMazemapLocationsParams{
			Limit:  req.Limit,
			Offset: req.Offset,
		})
	case "coords":
		locations, err = server.service.GetCoordsLocations(ctx, db.GetCoordsLocationsParams{
			Limit:  req.Limit,
			Offset: req.Offset,
		})
	case "address":
		locations, err = server.service.GetAddressLocations(ctx, db.GetAddressLocationsParams{
			Limit:  req.Limit,
			Offset: req.Offset,
		})
	case "none", "":
		locations, err = server.service.GetLocations(ctx, db.GetLocationsParams{
			Limit:  req.Limit,
			Offset: req.Offset,
		})
	}

	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, locations)
}

type getLocationRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// getLocation godoc
//
//	@Summary		Get location
//	@Description	Get a location by ID
//	@Tags			locations
//	@Produce		json
//	@Param			id	path		int	true	"Location ID"
//	@Success		200	{object}	db.Location
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/locations/{id} [get]
func (server *Server) getLocation(ctx *gin.Context) {
	var req getLocationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, fmt.Errorf("getLocation, ShouldBindQuery - %w", err))
		return
	}

	location, err := server.service.GetLocation(ctx, req.ID)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, location)
}

type createLocationRequest struct {
	NameNo string          `json:"name_no" binding:"required"`
	NameEn zero.String     `json:"name_en"`
	Type   db.LocationType `json:"type" binding:"required,locationtype"`

	MazemapCampusID zero.Int `json:"mazemap_campus_id"`
	MazemapPOIID    zero.Int `json:"mazemap_poi_id"`

	AddressStreet   zero.String `json:"address_street"`
	AddressPostcode zero.Int    `json:"address_postcode"`
	CityName        zero.String `json:"city_name"`

	CoordinateLat  zero.Float `json:"coordinate_lat"`
	CoordinateLong zero.Float `json:"coordinate_long"`

	URL zero.String `json:"url"`
}

// createLocation godoc
//
//	@Summary		Create location
//	@Description	Create a new location
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Param			location	body		createLocationRequest	true	"Location data"
//	@Success		200			{object}	db.Location
//	@Failure		400			{object}	errorResponse
//	@Failure		500			{object}	errorResponse
//	@Router			/locations [post]
func (server *Server) createLocation(ctx *gin.Context) {
	var req createLocationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[createLocationRequest](server, ctx, err)
		return
	}

	location, err := server.service.CreateLocation(ctx, db.CreateLocationParams{
		NameNo:          req.NameNo,
		NameEn:          req.NameEn,
		Type:            req.Type,
		AddressStreet:   req.AddressStreet,
		AddressPostcode: req.AddressPostcode,
		CityName:        req.CityName,
		CoordinateLat:   req.CoordinateLat,
		CoordinateLong:  req.CoordinateLong,
		Url:             req.URL,
		MazemapCampusID: req.MazemapCampusID,
		MazemapPoiID:    req.MazemapPOIID,
	})
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, location)
}

type updateLocationRequest struct {
	ID int32 `json:"id" binding:"required,min=1"`
	db.UpdateLocationParams
}

// updateLocation godoc
//
//	@Summary		Update location
//	@Description	Update a location by ID
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"Location ID"
//	@Param			request	body		db.UpdateLocationParams	true	"Location details"
//	@Success		200		{object}	db.Location
//	@Failure		400		{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Router			/locations [patch]
func (server *Server) updateLocation(ctx *gin.Context) {
	var req updateLocationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[updateLocationRequest](server, ctx, err)
		return
	}

	req.UpdateLocationParams.ID = req.ID

	location, err := server.service.UpdateLocation(ctx, req.UpdateLocationParams)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, location)
}

type deleteLocationRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// deleteLocation godoc
//
//	@Summary		Delete location
//	@Description	Delete a location by ID
//	@Tags			locations
//	@Produce		json
//	@Param			id	path		int	true	"Location ID"
//	@Success		200	{object}	db.Location
//	@Failure		400	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/locations/{id} [delete]
func (server *Server) deleteLocation(ctx *gin.Context) {
	var req deleteLocationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, fmt.Errorf("deleteLocation, ShouldBindUri - %w", err))
		return
	}

	location, err := server.service.SoftDeleteLocation(ctx, req.ID)
	err = db.ParseError(err)
	if err != nil {
		server.writeDBError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, location)
}
