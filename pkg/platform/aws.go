package platform

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	awsv2 "github.com/aws/aws-sdk-go-v2/aws"
	credentialsv2 "github.com/aws/aws-sdk-go-v2/credentials"
	dynamodbv2 "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func CreateAWSSession(accessKeyId, secretAccessKeyId, token, region string) (*session.Session, error) {
	cfg, err := createS3Credentials(accessKeyId, secretAccessKeyId, token, region)

	if err != nil {
		return nil, err
	}

	return session.NewSession(cfg)
}

func createS3Credentials(accessKeyId, secretAccessKeyId, token, region string) (*aws.Config, error) {
	creds := credentials.NewStaticCredentials(accessKeyId, secretAccessKeyId, token)
	_, err := creds.Get()

	if err != nil {
		// @todo handle error
		return nil, err
	}

	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)

	return cfg, nil
}

func CreateDynamodbConnection(session *session.Session) *dynamodb.DynamoDB {
	// dynamodbv2.New()
	// @todo change dynamodb.New to v2
	return dynamodb.New(session)
}

func Serve(handler interface{}) {
	lambda.Start(handler)
}
func CreateLocalClient() *dynamodbv2.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(awsv2.EndpointResolverFunc(
			func(service, region string) (awsv2.Endpoint, error) {
				return awsv2.Endpoint{URL: "http://localhost:8000"}, nil
			})),
		config.WithCredentialsProvider(credentialsv2.StaticCredentialsProvider{
			Value: awsv2.Credentials{
				AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	return dynamodbv2.NewFromConfig(cfg)
}
