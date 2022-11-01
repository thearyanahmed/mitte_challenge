package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/go-chi/chi/v5"

	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	routeHandler "github.com/thearyanahmed/mitte_challenge/pkg/handler"
	"github.com/thearyanahmed/mitte_challenge/pkg/platform"
)

var (
	db        dynamodbiface.DynamoDBAPI
	chiLambda *chiadapter.ChiLambda
)

// you can setup .env using the aws cli. eg:
// ```
// aws lambda update-function-configuration --function-name my-function \
// --environment "Variables={BUCKET=my-bucket,KEY=file.txt}"
// ```
// reference: https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html
func checkEnv() {
	keys := []string{"AWS_SECRET_ACCESS_KEY", "AWS_SECRET_KEY", "AWS_REGION"}

	for _, key := range keys {
		if os.Getenv(key) == "" {
			log.Fatalf("missing required key:%s\nyou can use export to set environment variables.\nusing .env in lambda will not satisfy requirements", key)
		}
	}
}

func init() {
	checkEnv()

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

	setupAdapter(r)
}

func main() {
	platform.Serve(handler)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return chiLambda.ProxyWithContext(ctx, req)
}

func setupAdapter(r *chi.Mux) {
	chiLambda = chiadapter.New(r)
}
