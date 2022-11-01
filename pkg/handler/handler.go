package handler

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/thearyanahmed/mitte_challenge/pkg/presenter"
)

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return presenter.BadRequest()
}
