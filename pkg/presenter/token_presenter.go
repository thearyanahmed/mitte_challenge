package presenter

import "github.com/thearyanahmed/mitte_challenge/pkg/entity"

type TokenResponse struct {
	Token string `json:"token"`
}

func FromToken(u entity.Token) TokenResponse {
	return TokenResponse{
		Token: u.Token,
	}
}
