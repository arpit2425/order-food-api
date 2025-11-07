package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"oilio.com/internal/model"
)

var validate = validator.New()

func ValidateOrderRequest(req model.OrderRequest) error {
	err := validate.Struct(req)
	if err == nil {
		return nil
	}

	var ve validator.ValidationErrors
	if errors, ok := err.(validator.ValidationErrors); ok {
		ve = errors
	}

	var errorMessages []string
	for _, fe := range ve {
		field := fe.Field()
		switch fe.Tag() {
		case "required":
			errorMessages = append(errorMessages, fmt.Sprintf("%s is required", field))
		case "min":
			errorMessages = append(errorMessages, fmt.Sprintf("%s must be at least %s", field, fe.Param()))
		case "max":
			errorMessages = append(errorMessages, fmt.Sprintf("%s must be at most %s", field, fe.Param()))
		default:
			errorMessages = append(errorMessages, fmt.Sprintf("%s is invalid", field))
		}
	}

	return fmt.Errorf("validation failed: %s", strings.Join(errorMessages, ", "))
}
