package serializer

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

type RequestValidatorInterface interface {
	Rules() govalidator.MapData
}

func ValidateJson(r *http.Request, validatable RequestValidatorInterface) url.Values {
	opts := govalidator.Options{
		Request: r,
		Data:    &validatable,
		Rules:   validatable.Rules(),
	}

	v := govalidator.New(opts)
	return v.ValidateJSON()
}