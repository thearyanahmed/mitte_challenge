package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

var (
	ctx = context.Background()
)

func TestUnregisteredRoute(t *testing.T) {
	req := events.APIGatewayProxyRequest{
		Path: "/user/create",
	}

	resp, err := handler(ctx, req)

	assert.IsType(t, err, nil)

	stringBody, _ := json.Marshal(map[string]string{"message": "/user/create"})
	assert.Equal(t, string(stringBody), string(resp.Body))
}
