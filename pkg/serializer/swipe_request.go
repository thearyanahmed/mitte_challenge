package serializer

import "github.com/thedevsaddam/govalidator"

type SwipeRequest struct {
	Preference     string `json:"preference"`
	ProfileOwnerID string `json:"profile_owner_id"`
}

func (r *SwipeRequest) Rules() govalidator.MapData {
	return govalidator.MapData{
		// @todo handle preference in a better way
		"preference": []string{"required","in:yes,no"},
		"profile_owner_id": []string{"required"},
	}
}

func (r *SwipeRequest) GetPreference() string {
	return r.Preference
}

func (r *SwipeRequest) GetProfileId() string {
	return r.ProfileOwnerID
}
