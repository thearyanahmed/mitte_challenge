package platform

import (
	"context"
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
