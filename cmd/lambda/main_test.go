package main

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

var (
	ctx     = context.Background()
	headers = map[string]string{"Content-Type": "application/json"}
)

func TestUnregisteredRoute(t *testing.T) {
	req := events.APIGatewayProxyRequest{
		Path:       "/user/create",
		Headers:    headers,
		HTTPMethod: http.MethodPost,
	}

	resp, err := handler(ctx, req)

	assert.IsType(t, err, nil)

	stringBody, _ := json.Marshal(map[string]string{"message": "/user/create"})
	assert.Equal(t, string(stringBody), string(resp.Body))
}
