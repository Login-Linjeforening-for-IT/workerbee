package internal

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound     = errors.New("could not find id")
	ErrInvalid  = errors.New("invalid user data")
	ErrUnauthorized = errors.New("unauthorized opperation")
	ErrorMap        = map[error]struct {
		Status  int
		Message string
	}{
		ErrNotFound:     {Status: http.StatusBadRequest, Message: "did not find document"},
		ErrInvalid:  {Status: http.StatusBadRequest, Message: "invalid user data"},
		ErrUnauthorized: {Status: http.StatusUnauthorized, Message: "unauthorized operation"},
	}
)

func HandleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	for k, v := range ErrorMap {
		if errors.Is(err, k) {
			c.JSON(v.Status, gin.H{"error": v.Message})
			return true
		}
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	return true
}
