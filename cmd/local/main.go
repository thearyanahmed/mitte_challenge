package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/joho/godotenv"

	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	routeHandler "github.com/thearyanahmed/mitte_challenge/pkg/handler"
	"github.com/thearyanahmed/mitte_challenge/pkg/platform"
)

var (
	db        dynamodbiface.DynamoDBAPI
	chiLambda *chiadapter.ChiLambda
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

	r := routeHandler.SetupRouter(&db)

	http.ListenAndServe(fmt.Sprintf("localhost:%s", getPort()), r)
}

func getPort() string {
	if os.Getenv("APP_PORT") == "" {
		return "8000"
	}

	return os.Getenv("APP_PORT")
}
