package validators

import (
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	instance *validator.Validate
	once     sync.Once
)

func Get() *validator.Validate {
	once.Do(func() {
		instance = validator.New()

		instance.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		_ = instance.RegisterValidation("cpf", func(fl validator.FieldLevel) bool {
			return isValidCPF(fl.Field().String())
		})
	})
	return instance
}
