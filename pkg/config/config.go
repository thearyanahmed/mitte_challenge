package config

import (
	"log"
	"os"
)

type EnvValues struct {
	DbUri, DbDatabase string
}

// CheckEnv
// you can set up .env using the aws cli. eg:
// ```
// aws lambda update-function-configuration --function-name my-function \
// --environment "Variables={BUCKET=my-bucket,KEY=file.txt}"
// ```
// reference: https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html
func CheckEnv() {
	keys := []string{"DB_URI", "DB_DATABASE"}

	for _, key := range keys {
		if os.Getenv(key) == "" {
			log.Fatalf("missing required key:%s\nyou can use export to set environment variables.\nusing .env in lambda will not satisfy requirements", key)
		}
	}
}

func GetEnvValues() EnvValues {
	return EnvValues{
		DbDatabase: os.Getenv("DB_DATABASE"),
		DbUri:      os.Getenv("DB_URI"),
	}
}
