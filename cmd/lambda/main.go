package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-chi/chi/v5"

	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/thearyanahmed/mitte_challenge/pkg/config"
	routeHandler "github.com/thearyanahmed/mitte_challenge/pkg/handler"
	"github.com/thearyanahmed/mitte_challenge/pkg/platform"
)

var (
	chiLambda *chiadapter.ChiLambda
)

func init() {
	envValues := config.GetEnvValues()

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
