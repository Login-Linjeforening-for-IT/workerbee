package internal

import "github.com/go-playground/validator/v10"

func LocalTimeAfterField(fl validator.FieldLevel) bool {
	// Get the field and the other field to compare
	field := fl.Field()
	param := fl.Param()

	// Get the other field by name
	otherField := fl.Parent().FieldByName(param)
	if !otherField.IsValid() {
		fl.Param()
		return false
	}

	t1, ok1 := field.Interface().(LocalTime)
	t2, ok2 := otherField.Interface().(LocalTime)

	if !ok1 || !ok2 {
		return false
	}

	// Check if Start < End
	return t1.Before(t2.Time)
}

func SetUpValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("beforeField", LocalTimeAfterField)
	return validate
}
