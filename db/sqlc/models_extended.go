package db

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

var validTimeTypeEnum validator.Func = func(fl validator.FieldLevel) bool {
	return TimeTypeEnum(fl.Field().String()).Valid()
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

var validJobType validator.Func = func(fl validator.FieldLevel) bool {
	return JobType(fl.Field().String()).Valid()
}

func BindJobTypeValidator(validator *validator.Validate, tag string) {
	validator.RegisterValidation(tag, validJobType)
}

func (nl NullJobType) MarshalJSON() ([]byte, error) {
	if nl.Valid {
		return json.Marshal(nl.JobType)
	} else {
		return json.Marshal(nil)
	}
}

func (nl *NullJobType) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null", `""`:
		nl.Valid = false
		return nil
	}

	var v JobType
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	nl.JobType = v
	nl.Valid = true

	return nil
}

var validLocationType validator.Func = func(fl validator.FieldLevel) bool {
	return LocationType(fl.Field().String()).Valid()
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
