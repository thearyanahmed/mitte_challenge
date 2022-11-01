package platform

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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
	return dynamodb.New(session)
}

func Serve(handler interface{}) {
	lambda.Start(handler)
}
