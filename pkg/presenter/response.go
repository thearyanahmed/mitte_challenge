package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func Response(statusCode int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = statusCode

	respBody, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	resp.Body = string(respBody)

	return &resp, err
}

func ErrBadRequest() (*events.APIGatewayProxyResponse, error) {
	msg := map[string]string{"message": "bad request."}

	return Response(http.StatusBadRequest, msg)
}

func ErrResponse(err error) (*events.APIGatewayProxyResponse, error) {
	msg := map[string]string{"message": err.Error()}
	return Response(http.StatusBadRequest, msg)
}
