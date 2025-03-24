// TODO: Add godoc

package api

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.login.no/tekkom/web/beehive/admin-api/images"
)

func (server *Server) handleImageUpload(ctx *gin.Context, dir string, ratioW int, ratioH int) {
	file, headers, err := ctx.Request.FormFile("file")
	if err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	err = images.CheckImage(file, headers.Size, ratioW, ratioH)
	if err != nil {
		server.writeError(ctx, http.StatusBadRequest, err)
		return
	}

	err = server.imageStore.UploadImage(dir, headers.Filename, file)
	if err != nil {
		switch err.(type) {
		case *images.DirNotFoundError:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}

//	@Summary		Post a event banner
//	@Description	Post a event banner
//	@Tags			images
//	@Produce		json
//	@Param			file	formData	file	true	"File"
//	@Router			/images/events/banner [post]
//
//	@Success		204	{object}	nil
func (server *Server) uploadEventImageBanner(ctx *gin.Context) {
	server.handleImageUpload(ctx, "img/events/banner/", images.BannerW, images.BannerH)
}

//	@Summary		Post a small event image
//	@Description	Post a small event image
//	@Tags			images
//	@Produce		json
//	@Param			file	formData	file	true	"File"
//	@Router			/images/events/small [post]
//
//	@Success		204	{object}	nil
func (server *Server) uploadEventImageSmall(ctx *gin.Context) {
	server.handleImageUpload(ctx, "img/events/small/", images.BannerW, images.BannerH)
}

//	@Summary		Post a job image
//	@Description	Post a job image
//	@Tags			images
//	@Produce		json
//	@Param			file	formData	file	true	"File"
//	@Router			/images/jobs [post]
//
//	@Success		204	{object}	nil
func (server *Server) uploadJobsImage(ctx *gin.Context) {
	server.handleImageUpload(ctx, "img/ads/", images.AdsW, images.AdsH)
}

//	@Summary		Post a organization image
//	@Description	Post a organization image
//	@Tags			images
//	@Produce		json
//	@Param			file	formData	file	true	"File"
//	@Router			/images/organizations [post]
//
//	@Success		204	{object}	nil
func (server *Server) uploadOrganizationImage(ctx *gin.Context) {
	server.handleImageUpload(ctx, "img/organizations/", images.OrgW, images.OrgH)
}

func (server *Server) handleImageList(ctx *gin.Context, dir string) {
	files, err := server.imageStore.GetImages(dir)
	if err != nil {
		switch err.(type) {
		case *images.DirNotFoundError:
			server.writeError(ctx, http.StatusNotFound, err)
		default:
			server.writeError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, files)
}

// TODO: Add godoc
//
//	@Summary		Get a list of events banner
//	@Description	Get a list of events banner
//	@Tags			images
//	@Produce		json
//	@Router			/images/events/banner [get]
//	@Success		200	{array}	images.FileDetails
func (server *Server) fetchEventsBannerList(ctx *gin.Context) {
	server.handleImageList(ctx, "img/events/banner/")
}

//	@Summary		Get a list of small event images
//	@Description	Get a list of small event images
//	@Tags			images
//	@Produce		json
//	@Router			/images/events/small [get]
//
//	@Success		200	{array}	images.FileDetails
func (server *Server) fetchEventsSmallList(ctx *gin.Context) {
	server.handleImageList(ctx, "img/events/small/")
}

//	@Summary		Get a list of job images
//	@Description	Get a list of job images
//	@Tags			images
//	@Produce		json
//	@Router			/images/jobs [get]
//
//	@Success		200	{array}	images.FileDetails
func (server *Server) fetchJobsList(ctx *gin.Context) {
	server.handleImageList(ctx, "img/ads/")
}

//	@Summary		Get a list of organization images
//	@Description	Get a list of organization images
//	@Tags			images
//	@Produce		json
//	@Router			/images/organizations [get]
//
//	@Success		200	{array}	images.FileDetails
func (server *Server) fetchOrganizationsList(ctx *gin.Context) {
	server.handleImageList(ctx, "img/organizations/")
}
