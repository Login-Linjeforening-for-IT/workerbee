package db

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

func IsValidTimeTypeEnum[T ~string](t T) bool {
	switch TimeTypeEnum(t) {
	case TimeTypeEnumDefault, TimeTypeEnumWholeDay, TimeTypeEnumNoEnd:
		return true
	default:
		return false
	}
}

var validTimeTypeEnum validator.Func = func(fl validator.FieldLevel) bool {
	return IsValidTimeTypeEnum(fl.Field().String())
}

func BindTimeTypeEnumValidator(validator *validator.Validate, tag string) {
	validator.RegisterValidation(tag, validTimeTypeEnum)
}

func (ns NullTimeTypeEnum) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.TimeTypeEnum)
	} else {
		return json.Marshal(nil)
	}
}

func (ns *NullTimeTypeEnum) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null", `""`:
		ns.Valid = false
		return nil
	}

	var v TimeTypeEnum
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	ns.TimeTypeEnum = v
	ns.Valid = true

	return nil
}

func IsValidLocationType[T ~string](t T) bool {
	switch LocationType(t) {
	case LocationTypeAddress, LocationTypeCoords, LocationTypeMazemap, LocationTypeNone:
		return true
	default:
		return false
	}
}

var validLocationType validator.Func = func(fl validator.FieldLevel) bool {
	return IsValidLocationType(fl.Field().String())
}

func BindLocationTypeValidator(validator *validator.Validate, tag string) {
	validator.RegisterValidation(tag, validLocationType)
}

func (nl NullLocationType) MarshalJSON() ([]byte, error) {
	if nl.Valid {
		return json.Marshal(nl.LocationType)
	} else {
		return json.Marshal(nil)
	}
}

func (nl *NullLocationType) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null", `""`:
		nl.Valid = false
		return nil
	}

	var v LocationType
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	nl.LocationType = v
	nl.Valid = true

	return nil
}
