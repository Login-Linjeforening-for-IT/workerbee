package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) authMiddleware(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("auth")
		if auth != secret {
			server.writeError(ctx, http.StatusUnauthorized, errors.New("invalid auth header"))
			ctx.Abort()
		}

		ctx.Next()
	}
}
