package main

// This file contains tests for running the application in lambda.

import (
	"context"
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
