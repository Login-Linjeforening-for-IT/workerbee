package api

import (
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

func (server *Server) uploadEventImageBanner(ctx *gin.Context) {
	server.uploadImageRequest(ctx, "img/events/banner/", 10, 4)
}

func (server *Server) uploadEventImageSmall(ctx *gin.Context) {
	server.uploadImageRequest(ctx, "img/events/small/", 10, 4)
}

func (server *Server) uploadJobsImage(ctx *gin.Context) {
	server.uploadImageRequest(ctx, "img/ads/", 3, 2)
}

func (server *Server) uploadOrganizationImage(ctx *gin.Context) {
	server.uploadImageRequest(ctx, "img/organizations/", 3, 2)
}

func (server *Server) fetchEventsBannerList(ctx *gin.Context) {
	server.fetchImageList(ctx, "img/events/banner/")
}

func (server *Server) fetchEventsSmallList(ctx *gin.Context) {
	server.fetchImageList(ctx, "img/events/small/")
}

func (server *Server) fetchJobsList(ctx *gin.Context) {
	server.fetchImageList(ctx, "img/ads/")
}

func (server *Server) fetchOrganizationsList(ctx *gin.Context) {
	server.fetchImageList(ctx, "img/organizations/")
}
