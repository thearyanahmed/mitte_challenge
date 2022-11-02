package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-chi/chi/v5"

	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	routeHandler "github.com/thearyanahmed/mitte_challenge/pkg/handler"
	"github.com/thearyanahmed/mitte_challenge/pkg/platform"
)

var (
	db        dynamodb.Client
	chiLambda *chiadapter.ChiLambda
)

// you can setup .env using the aws cli. eg:
// ```
// aws lambda update-function-configuration --function-name my-function \
// --environment "Variables={BUCKET=my-bucket,KEY=file.txt}"
// ```
// reference: https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html
func checkEnv() {
	keys := []string{"AWS_SECRET_ACCESS_KEY", "AWS_SECRET_KEY", "AWS_REGION", "DB_ENDPOINT"}

	for _, key := range keys {
		if os.Getenv(key) == "" {
			log.Fatalf("missing required key:%s\nyou can use export to set environment variables.\nusing .env in lambda will not satisfy requirements", key)
		}
	}
}

type EnvValues struct {
	AccessKey, SecretKey, Region, Token, DbEndpoint string
}

func getEnvValues() EnvValues {
	checkEnv()

	return EnvValues{
		AccessKey:  os.Getenv("AWS_SECRET_ACCESS_KEY"),
		SecretKey:  os.Getenv("AWS_SECRET_KEY"),
		Region:     os.Getenv("AWS_REGION"),
		Token:      "",
		DbEndpoint: os.Getenv("DB_ENDPOINT"),
	}
}

func init() {
	envValues := getEnvValues()

	db, err := platform.CreateDbConnection(envValues.AccessKey, envValues.SecretKey, envValues.Token, envValues.Region, envValues.DbEndpoint)

	if err != nil {
		log.Fatalf("could not establish connection to database.%v\n", err)
	}

	r := routeHandler.SetupRouter(db)

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
