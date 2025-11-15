package internal

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func LocalTimeBeforeField(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	otherField := fl.Parent().FieldByName(param)
	if !otherField.IsValid() {
		return false
	}

	var t1, t2 LocalTime

	if field.Kind() == reflect.Pointer {
		if field.IsNil() {
			return true
		}
		t1 = field.Elem().Interface().(LocalTime)
	} else {
		t1 = field.Interface().(LocalTime)
	}

	if otherField.Kind() == reflect.Pointer {
		if otherField.IsNil() {
			return true
		}
		t2 = otherField.Elem().Interface().(LocalTime)
	} else {
		t2 = otherField.Interface().(LocalTime)
	}

	// Save some context in Param
	// This can be used in HandleValidationError
	fl.Param()

	return t1.Before(t2)
}

func LocalTimeAfterField(fl validator.FieldLevel) bool {
	// Get the field and the other field to compare
	field := fl.Field()
	param := fl.Param() // normally the other field name

	otherField := fl.Parent().FieldByName(param)
	if !otherField.IsValid() {
		return false
	}

	var t1, t2 LocalTime

	if field.Kind() == reflect.Pointer {
		if field.IsNil() {
			return true
		}
		t1 = field.Elem().Interface().(LocalTime)
	} else {
		t1 = field.Interface().(LocalTime)
	}

	if otherField.Kind() == reflect.Pointer {
		if otherField.IsNil() {
			return true
		}
		t2 = otherField.Elem().Interface().(LocalTime)
	} else {
		t2 = otherField.Interface().(LocalTime)
	}

	// Save some context in Param
	// This can be used in HandleValidationError
	fl.Param()

	return t1.After(t2)
}

func SetUpValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("beforeField", LocalTimeBeforeField)
	validate.RegisterValidation("afterField", LocalTimeAfterField)
	return validate
}
