package serializer

import (
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
	"github.com/thedevsaddam/govalidator"
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

func ValidateRequest(r *http.Request, rules govalidator.MapData, data interface{}) url.Values {
	opts := govalidator.Options{
		Request: r,
		Data:    &data,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()

	return e
}
