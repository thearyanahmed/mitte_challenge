package handler

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/thearyanahmed/mitte_challenge/pkg/presenter"
)

type createUserHandler struct {
	// @todo use interface
	db *dynamodbiface.DynamoDBAPI
}

func NewCreateUserHandler(db *dynamodbiface.DynamoDBAPI) *createUserHandler {
	return &createUserHandler{
		db: db,
	}
}

func (h *createUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	presenter.Response(w, http.StatusOK, map[string]string{"message": "/user/create"})
}
