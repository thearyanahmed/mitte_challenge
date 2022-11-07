package main

import (
	"context"
	"fmt"
	"github.com/thearyanahmed/mitte_challenge/pkg/config"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/thearyanahmed/mitte_challenge/pkg/db"
	routeHandler "github.com/thearyanahmed/mitte_challenge/pkg/handler"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	envValues := config.GetEnvValues()

	client, database, err := db.ConnectToMongo(context.TODO(), envValues.DbUri, envValues.DbDatabase)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	aggregator := service.NewServiceAggregator(database)
	r := routeHandler.BootstrapRouter(aggregator)

	addr := fmt.Sprintf("0.0.0.0:%s", getPort())

	// @todo handle graceful shutdown
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
