package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PongHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
