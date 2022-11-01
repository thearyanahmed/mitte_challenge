package handler

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func handler() {
	fmt.Println("Handler")
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	// @todo : change here
	fmt.Println("method not allowed")
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = 405 // @todo use http.StatusCode

	stringBody, _ := json.Marshal(map[string]string{"message": "method not allowed."})
	resp.Body = string(stringBody)
	return &resp, nil
}
