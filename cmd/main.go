package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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
		fmt.Println("could not establish session, quiting", err)
		return
	} else {
		fmt.Println("Connected")
		return
	}

	db = dynamodb.New(awsSession)
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	default:
		return routeHandler.UnhandledMethod()
	}
}
