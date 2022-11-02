package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

var (
	ctx     = context.Background()
	headers = map[string]string{"Content-Type": "application/json"}
)

func TestCreateRandomUserRoute(t *testing.T) {
	req := events.APIGatewayProxyRequest{
		Path:       "/user/create",
		Headers:    headers,
		HTTPMethod: http.MethodPost,
	}

	resp, err := handler(ctx, req)

	assert.IsType(t, err, nil)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	expectedNonNullKeys := []string{"id", "name", "email", "password", "gender", "age"}
	for _, v := range expectedNonNullKeys {
		assert.Contains(t, resp.Body, v)
	}
}
