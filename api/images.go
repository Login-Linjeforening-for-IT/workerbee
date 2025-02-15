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

func (server *Server) uploadEventImageBanner(ctx *gin.Context) {
	server.handleImageUpload(ctx, "img/events/banner/", images.BannerW, images.BannerH)
}

func (server *Server) uploadEventImageSmall(ctx *gin.Context) {
	server.handleImageUpload(ctx, "img/events/small/", images.BannerW, images.BannerH)
}

func (server *Server) uploadJobsImage(ctx *gin.Context) {
	server.handleImageUpload(ctx, "img/ads/", images.AdsW, images.AdsH)
}

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
//	@Tags			jobs
//	@Produce		json
//	@Param			name	form		int	 true	"Name"
//	@Param			id		form		int	 true	"ID"
//	@Param			file	form		file true	"File"
//	@Router			/fetchEventBannerList [get]
func (server *Server) fetchEventsBannerList(ctx *gin.Context) {
	server.handleImageList(ctx, "img/events/banner/")
}

func (server *Server) fetchEventsSmallList(ctx *gin.Context) {
	server.handleImageList(ctx, "img/events/small/")
}

func (server *Server) fetchJobsList(ctx *gin.Context) {
	server.handleImageList(ctx, "img/ads/")
}

func (server *Server) fetchOrganizationsList(ctx *gin.Context) {
	server.handleImageList(ctx, "img/organizations/")
}
