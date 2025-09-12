package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler godoc
// @Summary      Ping endpoint
// @Description  Health check endpoint. Returns a simple pong message.
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /api/v2/ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
