package main

// This file contains tests for running the application in lambda.

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

var (
	ctx     = context.TODO()
	headers = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
)

type user struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
}

type loginValidationFailedResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details struct {
		Email []string `json:"email"`
	} `json:"details"`
}

type authToken struct {
	Token string `json:"token"`
}

type swipeResponse struct {
	Message         string `json:"message"`
	Matched         bool   `json:"preference_matched"`
	RecordedSwipeId string `json:"recorded_swipe_id"`
	MatchedSwipeId  string `json:"matched_swipe_id"`
}

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

func TestCreateRandomUser(t *testing.T) {
	// /user/create
	req := events.APIGatewayProxyRequest{
		Path:       "/user/create",
		Headers:    headers,
		HTTPMethod: http.MethodPost,
	}

	resp, err := handler(ctx, req)

	assert.IsType(t, err, nil)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	keys := []string{"id", "name", "email", "password", "gender", "age"}

	for _, k := range keys {
		assert.Contains(t, resp.Body, k)
	}
}

func createUser(t *testing.T) user {
	req := events.APIGatewayProxyRequest{
		Path:       "/user/create",
		Headers:    headers,
		HTTPMethod: http.MethodPost,
	}

	resp, err := handler(ctx, req)

	assert.Nil(t, err)

	data := user{}
	err = json.Unmarshal([]byte(resp.Body), &data)

	assert.Nil(t, err)

	return data
}

func TestUserCanLoginWithValidCredentials(t *testing.T) {
	user := createUser(t)

	body := url.Values{}
	body.Set("email", user.Email)
	body.Set("password", user.Password)

	req := events.APIGatewayProxyRequest{
		Path:       "/auth/login",
		Headers:    headers,
		HTTPMethod: http.MethodPost,
		Body:       body.Encode(),
	}

	resp, err := handler(ctx, req)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestAttemptToLoginWithInvalidDataReturnsError(t *testing.T) {
	body := url.Values{}
	body.Set("email", "some_random_user@that_should_not_exists.com")
	body.Set("password", "some_random_password")

	req := events.APIGatewayProxyRequest{
		Path:       "/auth/login",
		Headers:    headers,
		HTTPMethod: http.MethodPost,
		Body:       body.Encode(),
	}

	resp, err := handler(ctx, req)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var validationFailedResponse loginValidationFailedResponse

	err = json.Unmarshal([]byte(resp.Body), &validationFailedResponse)

	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, validationFailedResponse.Code)
	assert.Equal(t, validationFailedResponse.Message, "validation failed.")
	assert.Equal(t, validationFailedResponse.Details.Email[0], "The email field must be a valid email address")
}

func login(email, password string) (events.APIGatewayProxyResponse, error) {
	body := url.Values{}
	body.Set("email", email)
	body.Set("password", password)

	req := events.APIGatewayProxyRequest{
		Path:       "/auth/login",
		Headers:    headers,
		HTTPMethod: http.MethodPost,
		Body:       body.Encode(),
	}

	return handler(ctx, req)
}

func loginWithNewlyCreatedUser(t *testing.T) (user, string, events.APIGatewayProxyResponse) {
	user := createUser(t)

	resp, err := login(user.Email, user.Password)

	assert.Nil(t, err)

	var authToken authToken

	err = json.Unmarshal([]byte(resp.Body), &authToken)

	assert.Nil(t, err)

	return user, authToken.Token, resp
}
