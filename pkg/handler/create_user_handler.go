package handler

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func CreateUser(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}

	data := req.Body

	stringBody, _ := json.Marshal(map[string]string{"message": "method is allowed.", "data": data})
	resp.Body = string(stringBody)

	return &resp, nil
}
