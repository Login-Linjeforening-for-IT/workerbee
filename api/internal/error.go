package internal

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	ErrNotFound               = errors.New("could not find id")
	ErrNoRow                  = errors.New("no row found")
	ErrInvalid                = errors.New("invalid user data")
	ErrInvalidImagePath       = errors.New("invalid image path")
	ErrImageTooLarge          = errors.New("image size exceeds maximum limit")
	ErrInvalidImageRatio      = errors.New("invalid image aspect ratio")
	ErrUnauthorized           = errors.New("unauthorized opperation")
	ErrInvalidForeignKey      = errors.New("error foreign key does not exist")
	ErrInvalidAudience        = errors.New("invalid audience does not exist in enum")
	ErrInvalidTimeType        = errors.New("invalid time type does not exist in enum")
	ErrInvalidCategory        = errors.New("invalid category does not exist in enum")
	ErrInvalidLocationType    = errors.New("invalid location type does not exist in enum")
	ErrInvalidJobType         = errors.New("invalid job type does not exist in enum")
	ErrS3ClientNotInitialized = errors.New("s3 client not initialized")
	ErrorMap                  = map[error]struct {
		Status  int
		Message string
	}{
		ErrNotFound:          {Status: http.StatusBadRequest, Message: "did not find document"},
		ErrNoRow:             {Status: http.StatusBadRequest, Message: "no row found"},
		ErrInvalid:           {Status: http.StatusBadRequest, Message: "invalid user data"},
		ErrInvalidImagePath:  {Status: http.StatusBadRequest, Message: "invalid image path"},
		ErrUnauthorized:      {Status: http.StatusUnauthorized, Message: "unauthorized operation"},
		ErrInvalidForeignKey: {Status: http.StatusBadRequest, Message: "error foreign key does not exist"},
		ErrInvalidAudience:   {Status: http.StatusBadRequest, Message: "invalid audience does not exist in enum"},
		ErrInvalidTimeType:   {Status: http.StatusBadRequest, Message: "invalid time type does not exist in enum"},
		ErrInvalidCategory:   {Status: http.StatusBadRequest, Message: "invalid category does not exist in enum"},
		ErrInvalidLocationType: {
			Status:  http.StatusBadRequest,
			Message: "invalid location type does not exist in enum",
		},
		ErrInvalidJobType:         {Status: http.StatusBadRequest, Message: "invalid job type does not exist in enum"},
		ErrImageTooLarge:          {Status: http.StatusBadRequest, Message: "image size exceeds maximum limit, max 1MB"},
		ErrInvalidImageRatio:      {Status: http.StatusBadRequest, Message: "invalid image aspect ratio, max 2.5"},
		ErrS3ClientNotInitialized: {Status: http.StatusInternalServerError, Message: "s3 client not initialized"},
	}
)

func HandleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	for k, v := range ErrorMap {
		if errors.Is(err, k) {
			c.JSON(v.Status, gin.H{"error": v.Message})
			log.Println("Got error: ", err)
			return true
		}
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	log.Println("Got error: ", err)
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
