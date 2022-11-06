package serializer

import (
	"net/http"
	"net/url"

	saddamValidator "github.com/thedevsaddam/govalidator"
	"go.mongodb.org/mongo-driver/bson"
)

// ProfileFilterRequest represents the request for create user.
type ProfileFilterRequest struct {
	Age    string `json:"age" validate:"numeric,omitempty"`
	Gender string `json:"gender" validate:"omitempty,oneof=male female"`
}

func ValidateRequest(r *http.Request, requestStruct interface{}) url.Values {
	rules := saddamValidator.MapData{
		"gender": []string{"in:male,female"},
		"age":    []string{"numeric"},
	}

	opts := saddamValidator.Options{
		Request: r,
		Data:    &requestStruct,
		Rules:   rules,
	}

	v := saddamValidator.New(opts)
	e := v.ValidateJSON()

	return e
}

// Bind only implements interface contract.
func (u *ProfileFilterRequest) Bind(r *http.Request) error {
	return nil
}

func (u *ProfileFilterRequest) ToKeyValuePair() map[string]string {
	mapped := make(map[string]string)

	if u.Age != "" {
		// @todo parse int
		mapped["age"] = u.Age
	}

	if u.Gender != "" {
		mapped["gender"] = u.Gender
	}

	return mapped
}

func (u *ProfileFilterRequest) ToX() bson.D {
	var filters bson.D

	if u.Age != "" {
		filters = append(filters, bsonFilter("age", u.Age)...)
	}
	if u.Gender != "" {
		filters = append(filters, bsonFilter("gender", u.Gender)...)
	}

	return filters
}

func bsonFilter(key string, value string) bson.D {
	return bson.D{{Key: key, Value: value}}
}
