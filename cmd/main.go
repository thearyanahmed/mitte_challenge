package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/joho/godotenv"

	routeHandler "github.com/thearyanahmed/mitte_challenge/pkg/handler"
	"github.com/thearyanahmed/mitte_challenge/pkg/platform"
)

var (
	db dynamodbiface.DynamoDBAPI
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	awsSession, err := platform.CreateAWSSession(
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		os.Getenv("AWS_SECRET_KEY"),
		"",
		os.Getenv("AWS_REGION"),
	)

	if err != nil {
		// @todo handle better way, make sure to log
		log.Fatalf("could not establish session. %v\n", err)
		return
	}

	db = platform.CreateDynamodbConnection(awsSession)

	platform.Serve(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// @todo also make sure we are handling httpMethod
	switch req.Path {
	case routeHandler.CreateUserRoutePath: // register route here
		return routeHandler.CreateUser(req)
	default:
		return routeHandler.UnhandledMethod()
	}
}
