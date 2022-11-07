package main

import (
	"context"
	"github.com/thearyanahmed/mitte_challenge/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-chi/chi/v5"

	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	routeHandler "github.com/thearyanahmed/mitte_challenge/pkg/handler"
	"github.com/thearyanahmed/mitte_challenge/pkg/platform"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
)

var (
	chiLambda *chiadapter.ChiLambda
	client    *mongo.Client
)

func init() {
	config.CheckEnv()

	envValues := config.GetEnvValues()

	dbc, database, err := platform.ConnectToMongo(context.TODO(), envValues.DbUri, envValues.DbDatabase)
	client = dbc

	if err != nil {
		log.Fatal(err)
	}

	aggregator := service.NewServiceAggregator(database)
	r := routeHandler.BootstrapRouter(aggregator)

	setupAdapter(r)
}

func main() {
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	platform.Serve(handler)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return chiLambda.ProxyWithContext(ctx, req)
}

func setupAdapter(r *chi.Mux) {
	chiLambda = chiadapter.New(r)
}
