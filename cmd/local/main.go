package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/thearyanahmed/mitte_challenge/pkg/config"
	routeHandler "github.com/thearyanahmed/mitte_challenge/pkg/handler"
	"github.com/thearyanahmed/mitte_challenge/pkg/platform"
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

	r := routeHandler.SetupRouter(db)

	http.ListenAndServe(fmt.Sprintf("localhost:%s", getPort()), r)
}

func getPort() string {
	if os.Getenv("APP_PORT") == "" {
		return "8000"
	}

	return os.Getenv("APP_PORT")
}
