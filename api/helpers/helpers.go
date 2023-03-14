package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

func printError(err error) string{
	errorString := fmt.Sprintf("Request Error: %v",err.Error())

	fmt.Println(errorString)

	return errorString
}

func stringFromJson(obj interface{}) string{
	jsonString, err := json.MarshalIndent(obj, "", "    ");

	if err != nil{
		return "Error json indent"
	}

	return string(jsonString) 
}

func BindJsonOrAbort (obj interface{}, c *gin.Context) ([]ErrorMsg) {

	if err := c.BindJSON(obj); err != nil {
		fieldErrors := GetErrors(err)
		printError(errors.New(stringFromJson(fieldErrors)))


		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fieldErrors,
		})

		return fieldErrors
	}

	return nil;
}
