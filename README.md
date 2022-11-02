// Write about branching

Useful commands :

export $(grep -v '^#' .env | xargs) && aws dynamodb create-table --cli-input-json file://deployments/schema/users.json

export $(grep -v '^#' .env | xargs) && aws dynamodb list-tables --endpoint-url http://localhost:8001


Working /
aws dynamodb create-table --cli-input-json file://deployments/schema/users.json --endpoint-url http://localhost:8001


aws dynamodb list-tables --endpoint-url http://localhost:8001
