package serializer

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

// NewValidator returns a new validator.
func NewValidator() *validator.Validate {
	validator := validator.New()
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})
	return validator
}
