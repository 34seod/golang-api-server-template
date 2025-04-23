package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// =================================================================
type APIResponse struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}

func response(c *gin.Context, statusCode int, data any, err any) {
	switch {
	case statusCode == http.StatusNoContent:
		c.Status(http.StatusNoContent)
	case statusCode >= 200 && statusCode < 300:
		c.SecureJSON(statusCode, APIResponse{Data: data})
	default:
		c.SecureJSON(statusCode, APIResponse{Error: err})
	}
}

// =================================================================
func errorCtrl(status int, err error, msg ...any) (int, any) {
	log.Print("\033[31m[ERROR] ", err.Error())
	if len(msg) > 0 {
		return status, msg[0]
	}
	return status, err.Error()
}

// =================================================================
type validateError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func validateErrors(err error) []validateError {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		out := make([]validateError, len(validationErrors))
		for i, fieldError := range validationErrors {
			out[i] = validateError{strings.ToLower(fieldError.Field()), msgForTag(fieldError)}
		}
		return out
	}

	return nil
}

func msgForTag(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "This field is in email format only"
	case "len":
		return fmt.Sprintf("This field needs %s digit", fieldError.Param())
	}
	return fieldError.Error()
}

// =================================================================
