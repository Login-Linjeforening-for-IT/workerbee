package api

import (
	"database/sql"
	"net/http"

	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null/zero"
)

type getLocationsRequest struct {
	Limit  int32  `form:"limit,default=20"`
	Offset int32  `form:"offset"`
	Type   string `form:"type"`
}

func (server *Server) getLocations(ctx *gin.Context) {
	var req getLocationsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
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

	if err != nil {
		server.writeError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, locations)
}

type getLocationRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getLocation(ctx *gin.Context) {
	var req getLocationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	location, err := server.service.GetLocation(ctx, req.ID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
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
	if err != nil {
		server.writeError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, location)
}

type updateLocationRequest struct {
	ID int32 `json:"id" binding:"required,min=1"`
	db.UpdateLocationParams
}

func (server *Server) updateLocation(ctx *gin.Context) {
	var req updateLocationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeValidationError[updateLocationRequest](server, ctx, err)
		return
	}

	req.UpdateLocationParams.ID = req.ID

	location, err := server.service.UpdateLocation(ctx, req.UpdateLocationParams)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, location)
}
