package common

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	FailedField string
	Tag         string
	Value       string
}

type ErrorResponse struct {
	Errors []*ValidationError `json:"errors"`
}

func (err *ErrorResponse) ToJson() string {
	if err.Errors != nil {
		data, _ := json.Marshal(err.Errors)
		return string(data)
	} else {
		return ""
	}
}

func Validator(structInstance interface{}) ErrorResponse {
	var validate = validator.New()
	var errors []*ValidationError
	err := validate.Struct(structInstance)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := ValidationError{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			}
			errors = append(errors, &element)
		}
	}
	return ErrorResponse{Errors: errors}
}
