// Write about branching

Useful commands :

export $(grep -v '^#' .env | xargs) && aws dynamodb create-table --cli-input-json file://deployments/schema/users.json

export $(grep -v '^#' .env | xargs) && aws dynamodb list-tables --endpoint-url http://localhost:8001


Working /
aws dynamodb create-table --cli-input-json file://deployments/schema/users.json --endpoint-url http://localhost:8001


aws dynamodb list-tables --endpoint-url http://localhost:8001


curl --request GET \
  --url http://127.0.0.1:8080/profile \
  --header 'Content-Type: application/json' \
  --header 'Authorization: 7259E42039C9947C8DBD664EEF3610C27C61D5B366D441BE3145460B76024E64' \
  --data '{
	"age": "10",
    "gender": "male"
}'