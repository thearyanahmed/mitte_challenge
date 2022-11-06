package serializer

import (
	"github.com/thedevsaddam/govalidator"
)

type ProfileFilterRequest struct {
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

func (r *ProfileFilterRequest) Rules() govalidator.MapData {
	return govalidator.MapData{
		"gender": []string{"in:male,female"},
		"age":    []string{"numeric","min:10","max:100"},
	}
}

func (r *ProfileFilterRequest) ToKeyValuePair() map[string]string {
	kvMap := map[string]string{}

	if r.Age != "" {
		kvMap["age"] = r.Age
	}

	if r.Gender != "" {
		kvMap["gender"] = r.Age
	}

	return kvMap
}