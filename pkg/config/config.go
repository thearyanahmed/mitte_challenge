package config

import (
	"log"
	"os"
)

type EnvValues struct {
	AccessKey, SecretKey, Region, Token, DbEndpoint string
}

// you can setup .env using the aws cli. eg:
// ```
// aws lambda update-function-configuration --function-name my-function \
// --environment "Variables={BUCKET=my-bucket,KEY=file.txt}"
// ```
// reference: https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html
func checkEnv() {
	keys := []string{"AWS_SECRET_ACCESS_KEY", "AWS_SECRET_KEY", "AWS_REGION", "DB_ENDPOINT"}

	for _, key := range keys {
		if os.Getenv(key) == "" {
			log.Fatalf("missing required key:%s\nyou can use export to set environment variables.\nusing .env in lambda will not satisfy requirements", key)
		}
	}
}

func GetEnvValues() EnvValues {
	checkEnv()

	return EnvValues{
		AccessKey:  os.Getenv("AWS_SECRET_ACCESS_KEY"),
		SecretKey:  os.Getenv("AWS_SECRET_KEY"),
		Region:     os.Getenv("AWS_REGION"),
		Token:      "",
		DbEndpoint: os.Getenv("DB_ENDPOINT"),
	}
}
