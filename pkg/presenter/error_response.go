package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func ErrorResponse(statusCode int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = statusCode

	respBody, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	resp.Body = string(respBody)

	return &resp, err
}

func BadRequest() (*events.APIGatewayProxyResponse, error) {
	msg := map[string]string{"message": "bad request."}

	return ErrorResponse(http.StatusBadRequest, msg)
}
