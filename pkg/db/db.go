package db

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Serve(handler interface{}) {
	lambda.Start(handler)
}

func ConnectToMongo(ctx context.Context, uri, db string) (*mongo.Client, *mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return nil, nil, err
	}

	database := client.Database(db)

	return client, database, nil
}
