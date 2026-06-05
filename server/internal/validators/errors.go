package validators

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatErrors(err error) map[string]string {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return map[string]string{"error": "invalid input"}
	}

	out := make(map[string]string, len(ve))
	for _, fe := range ve {
		out[fe.Field()] = messageFor(fe)
	}
	return out
}

func messageFor(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "required field"
	case "email":
		return "invalid email"
	case "cpf":
		return "invalid cpf"
	case "min":
		return fmt.Sprintf("minimum %s characters", fe.Param())
	case "max":
		return fmt.Sprintf("maximum %s characters", fe.Param())
	case "len":
		return fmt.Sprintf("must be exactly %s characters", fe.Param())
	case "gte":
		return fmt.Sprintf("must be greater than or equal to %s", fe.Param())
	case "gt":
		return "must be greater than zero"
	case "boolean":
		return "must be a boolean field"
	default:
		return "invalid value"
	}
}
