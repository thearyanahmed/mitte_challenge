package serializer

import (
	"github.com/gorilla/schema"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

type RequestValidatorInterface interface {
	Rules() govalidator.MapData
}

// ValidatePostForm @todo could be improved separating concerns
func ValidatePostForm(r *http.Request, validatable RequestValidatorInterface) url.Values {
	opts := govalidator.Options{
		Request: r,
		Data:    &validatable,
		Rules:   validatable.Rules(),
	}

	err := r.ParseForm()

	if err != nil {
		return url.Values{
			"form_parse_error": []string{err.Error()},
		}
	}

	// r.PostForm is a map of our POST form values
	err = schema.NewDecoder().Decode(validatable, r.PostForm)
	if err != nil {
		return url.Values{
			"decoder_error": []string{err.Error()},
		}
	}

	v := govalidator.New(opts)
	return v.Validate()
}

func ValidateGetQuery(r *http.Request, validatable RequestValidatorInterface) url.Values {
	opts := govalidator.Options{
		Request: r,
		Data:    &validatable,
		Rules:   validatable.Rules(),
	}

	v := govalidator.New(opts)
	return v.Validate()
}
