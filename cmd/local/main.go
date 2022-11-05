package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/thearyanahmed/mitte_challenge/pkg/config"
	routeHandler "github.com/thearyanahmed/mitte_challenge/pkg/handler"
	"github.com/thearyanahmed/mitte_challenge/pkg/platform"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
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

	// todo add concurrency, use wait groups
	// make sure to add them inside lambda as well
	if err = platform.WaitForDB(context.Background(), db, "users"); err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf("0.0.0.0:%s", getPort())

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
