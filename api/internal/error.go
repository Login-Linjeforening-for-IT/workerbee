package internal

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidSort = errors.New("invalid sorting data")
	ErrInvalidId = errors.New("invalid id")
)

func HandleError(c *gin.Context, err error, msg string, status int) {
	log.Println(msg+": ", err)
	c.AbortWithStatus(status)
}
