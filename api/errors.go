package api

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type errorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
	Type   string `json:"type"`
}

func newErrorResponse(status int, err error) errorResponse {
	res := errorResponse{
		Status: status,
	}

	if err != nil {
		res.Error = err.Error()
		res.Type = reflect.TypeOf(err).String()
	}

	return res
}

type RedactedError struct {
	ID      string `json:"id"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *RedactedError) Error() string {
	return fmt.Sprintf("(id=%s) %s", e.ID, e.Message)
}

func (server *Server) writeError(ctx *gin.Context, status int, err error) {
	if status >= 500 {
		id := uuid.NewString()
		err = &RedactedError{
			ID:      id,
			Status:  status,
			Message: "Something went wrong. Contact admin if problem persists.",
		}

		server.logger.Error().Err(err).Str("error-id", id).Send()
	}

	ctx.JSON(status, newErrorResponse(status, err))
}
