package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
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

	err = waitForTable(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf("0.0.0.0:%s", getPort())

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("could not serve. %v\n", err)
	}
}

func waitForTable(ctx context.Context, db *dynamodb.Client) error {
	w := dynamodb.NewTableExistsWaiter(db)
	err := w.Wait(ctx,
		&dynamodb.DescribeTableInput{
			TableName: aws.String("users"),
		},
		20*time.Second,
		func(o *dynamodb.TableExistsWaiterOptions) {
			o.MaxDelay = 5 * time.Second
			o.MinDelay = 5 * time.Second
		})

	return err
}

func getPort() string {
	if os.Getenv("APP_PORT") == "" {
		return "8080"
	}

	return os.Getenv("APP_PORT")
}
