package common

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
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
			//fmt.Println("Tag", err.Tag())
			//fmt.Println("ActualTag", err.ActualTag())
			//fmt.Println("Field", err.Field())
			//fmt.Println("Error", err.Error())
			//fmt.Println("Param", err.Param())
			//fmt.Println("Value", err.Value())
			//fmt.Println("Kind", err.Kind())
			//fmt.Println("Namespace", err.Namespace())
			//fmt.Println("StructField", err.StructField())
			//fmt.Println("StructNamespace", err.StructNamespace())
			//fmt.Println("Type", err.Type())

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
