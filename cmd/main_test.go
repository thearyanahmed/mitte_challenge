package main

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

// func TestHandler(t *testing.T) {
// 	tests := []struct {
// 		request events.APIGatewayProxyRequest
// 		expect  string
// 		err     error
// 	}{
// 		{
// 			// Test that the handler responds with the correct response
// 			// when a valid name is provided in the HTTP body
// 			request: events.APIGatewayProxyRequest{
// 				Body: "Paul",
// 				Path: "/user/create",
// 			},
// 			expect: "Hello Paul",
// 			err:    nil,
// 		},
// 	}

// 	for _, test := range tests {
// 		response, err := handler(test.request)
// 		assert.IsType(t, test.err, err)
// 		assert.Equal(t, test.expect, response.Body)
// 	}
// }

func TestUnregisteredRoute(t *testing.T) {
	req := events.APIGatewayProxyRequest{
		Path: "/random/endpoint/that/should/return/error",
	}

	resp, err := handler(req)

	assert.IsType(t, err, nil)

	stringBody, _ := json.Marshal(map[string]string{"message": "bad request."})
	assert.Equal(t, string(stringBody), resp.Body)
}
