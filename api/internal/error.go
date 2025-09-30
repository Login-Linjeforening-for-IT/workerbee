package internal

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	ErrNotFound          = errors.New("could not find id")
	ErrInvalid           = errors.New("invalid user data")
	ErrUnauthorized      = errors.New("unauthorized opperation")
	ErrInvalidForeignKey = errors.New("error foreign key does not exist")
	ErrorMap             = map[error]struct {
		Status  int
		Message string
	}{
		ErrNotFound:          {Status: http.StatusBadRequest, Message: "did not find document"},
		ErrInvalid:           {Status: http.StatusBadRequest, Message: "invalid user data"},
		ErrUnauthorized:      {Status: http.StatusUnauthorized, Message: "unauthorized operation"},
		ErrInvalidForeignKey: {Status: http.StatusBadRequest, Message: "error foreign key does not exist"},
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

func HandleValidationError[T any](c *gin.Context, body T, validate validator.Validate) bool {
	if err := validate.Struct(body); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			validationMessages := make([]string, 0, len(errs))
			eventType := reflect.TypeOf(body)
			if eventType.Kind() == reflect.Ptr {
				eventType = eventType.Elem()
			}

			for _, e := range errs {
				fieldName := e.Field()
				if f, found := eventType.FieldByName(e.StructField()); found {
					if jsonTag := f.Tag.Get("json"); jsonTag != "" {
						fieldName = strings.Split(jsonTag, ",")[0]
					}
				}
				validationMessages = append(validationMessages, fmt.Sprintf("%s failed on %s validation", fieldName, e.Tag()))
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationMessages,
			})
			return true
		}
		HandleError(c, err)
		return true
	}
	return false
}
