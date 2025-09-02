package internal

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, msg string, status int) {
	log.Println(msg+": ", err)
	c.AbortWithStatus(status)
}
