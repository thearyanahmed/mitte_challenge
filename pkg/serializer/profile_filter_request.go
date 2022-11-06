package serializer

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

type ProfileFilterRequest struct {
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

func (r *ProfileFilterRequest) Rules() govalidator.MapData {
	return govalidator.MapData{
		"age":    []string{"numeric"},
		"gender": []string{"in:male,female"},
	}
}

func NewProfileFilterRequestFromQuery(r *http.Request) *ProfileFilterRequest {
	filterFromQuery := NewProfileFilterRequest("", "")
	filterFromQuery.PopulateUsingQuery(r)

	return filterFromQuery
}

func NewProfileFilterRequest(age, gender string) *ProfileFilterRequest {
	return &ProfileFilterRequest{
		Age:    age,
		Gender: gender,
	}
}

func (r *ProfileFilterRequest) PopulateUsingQuery(req *http.Request) {
	if req.URL.Query().Get("age") != "" {
		r.Age = req.URL.Query().Get("age")
	}

	if req.URL.Query().Get("gender") != "" {
		r.Gender = req.URL.Query().Get("gender")
	}
}

func (r *ProfileFilterRequest) ToKeyValuePair() map[string]string {
	kvMap := map[string]string{}

	if r.Age != "" {
		kvMap["age"] = r.Age
	}

	if r.Gender != "" {
		kvMap["gender"] = r.Gender
	}

	return kvMap
}
