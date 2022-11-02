package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/thearyanahmed/mitte_challenge/pkg/config"
	routeHandler "github.com/thearyanahmed/mitte_challenge/pkg/handler"
	"github.com/thearyanahmed/mitte_challenge/pkg/platform"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
)

var (
	chiLambda *chiadapter.ChiLambda
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	envValues := config.GetEnvValues()
	db, err := platform.CreateDbConnection(envValues.AccessKey, envValues.SecretKey, envValues.Token, envValues.Region, envValues.DbEndpoint)

	if err != nil {
		log.Fatal(err)
	}

	aggregator := service.NewServiceAggregator(db)
	r := routeHandler.SetupRouter(aggregator)

	err = platform.WaitForDB(context.Background(), db, "users")
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf("0.0.0.0:%s", getPort())

	fmt.Printf("connected to db. willbe serving on %s inside container.", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("could not serve. %v\n", err)
	}
}

func getPort() string {
	if os.Getenv("APP_PORT") == "" {
		return "8080"
	}

	return os.Getenv("APP_PORT")
}
