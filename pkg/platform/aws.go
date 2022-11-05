package platform

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/aws"

	awsv2 "github.com/aws/aws-sdk-go-v2/aws"
	credentialsv2 "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Serve(handler interface{}) {
	lambda.Start(handler)
}

func ConnectToMongo() (*mongo.Client, *mongo.Database, error){
	// @todo make it dynamic, use .env?
	uri := "mongodb://db:27017"
	db := "app_2"

	// @todo context.Todo()?
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		fmt.Println("ERRRROR",err)
		return nil, nil, err
	}

	database := client.Database(db)

	return client, database, nil
}

func CreateDbConnection(accessKeyId, secretAccessKeyId, token, region, endpoint string) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithEndpointResolver(awsv2.EndpointResolverFunc(
			func(service, region string) (awsv2.Endpoint, error) {
				return awsv2.Endpoint{URL: endpoint}, nil
			})),
		config.WithCredentialsProvider(credentialsv2.StaticCredentialsProvider{
			Value: awsv2.Credentials{
				AccessKeyID: accessKeyId, SecretAccessKey: secretAccessKeyId, SessionToken: token,
			},
		}),
	)
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}

func WaitForDB(ctx context.Context, db *dynamodb.Client, table string) error {
	w := dynamodb.NewTableExistsWaiter(db)
	err := w.Wait(ctx,
		&dynamodb.DescribeTableInput{
			TableName: aws.String(table),
		},
		5*time.Second,
		func(o *dynamodb.TableExistsWaiterOptions) {
			o.MaxDelay = 5 * time.Second
			o.MinDelay = 5 * time.Second
		})

	return err
}
