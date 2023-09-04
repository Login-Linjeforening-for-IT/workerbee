package api

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

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

		errChain := []string{}
		for chainErr := err; chainErr != nil; chainErr = errors.Unwrap(chainErr) {
			errChain = append(errChain, chainErr.Error())
		}

		server.logger.Error().Err(err).Str("error-id", id).Str("error-chain", strings.Join(errChain, " -- ")).Int("error-chain-length", len(errChain)).Send()

		err = &RedactedError{
			ID:      id,
			Status:  status,
			Message: "Something went wrong. Contact admin if problem persists.",
		}
	}

	ctx.JSON(status, newErrorResponse(status, err))
}
