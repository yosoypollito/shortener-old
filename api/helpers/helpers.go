package helpers

import (
	"github.com/go-playground/validator/v10"
	"errors"
)

type ErrorMsg struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "email":
		return "Should be a valid email"
	case "url":
		return "Should be a valid url"
	}
	return "Unkown Error"
}

func GetErrors(err error) []ErrorMsg {
	var ve validator.ValidationErrors

	var out []ErrorMsg

	if errors.As(err, &ve) {
		for _, fe := range ve {
			out = append(out, ErrorMsg{fe.Field(), getErrorMsg(fe)})
		}
	}
	return out
} 
