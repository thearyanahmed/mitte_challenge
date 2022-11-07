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

func TestHealthCheckRoute(t *testing.T) {
	req := events.APIGatewayProxyRequest{
		Path:       "/health-check",
		Headers:    headers,
		HTTPMethod: http.MethodGet,
	}

	resp, err := handler(ctx, req)

	assert.IsType(t, err, nil)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	expectedNonNullKeys := []string{"status", "time"}
	for _, v := range expectedNonNullKeys {
		assert.Contains(t, resp.Body, v)
	}
}
