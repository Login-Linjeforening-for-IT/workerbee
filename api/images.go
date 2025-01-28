package api

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"

	"gitlab.login.no/tekkom/web/beehive/admin-api/images"
	"github.com/gin-gonic/gin"
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

	// Use the name from the form if it exists, otherwise use the filename
	name := ctx.Request.FormValue("name")
	if name == "" {
		name = headers.Filename
	}

	// err = server.imageStore.UploadImage(dir, name, file)
	// if err != nil {
	// 	switch err.(type) {
	// 	case *images.DirNotFoundError:
	// 		server.writeError(ctx, http.StatusNotFound, err)
	// 	default:
	// 		server.writeError(ctx, http.StatusInternalServerError, err)
	// 	}
	// 	return
	// }

	ctx.Status(http.StatusNoContent)
}

func (server *Server) uploadEventImageBanner(ctx *gin.Context) {
	server.handleImageUpload(ctx, "img/events/banner/", 10, 4)
}

func (server *Server) uploadEventImageSmall(ctx *gin.Context) {
	server.handleImageUpload(ctx, "img/events/small/", 10, 4)
}

func (server *Server) uploadJobsImage(ctx *gin.Context) {
	server.handleImageUpload(ctx, "img/ads/", 3, 2)
}

func (server *Server) uploadOrganizationImage(ctx *gin.Context) {
	server.handleImageUpload(ctx, "img/organizations/", 3, 2)
}

func (server *Server) handleImageList(ctx *gin.Context, dir string) {
	// files, err := server.imageStore.GetImages(dir)
	// if err != nil {
	// 	switch err.(type) {
	// 	case *images.DirNotFoundError:
	// 		server.writeError(ctx, http.StatusNotFound, err)
	// 	default:
	// 		server.writeError(ctx, http.StatusInternalServerError, err)
	// 	}
	// 	return
	// }

	// ctx.JSON(http.StatusOK, files)
}

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
