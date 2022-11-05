package schema

type UserTraitSchema struct {
	ID string `json:"id" dynamodbav:"id"`
	Value int8 `json:"value"  dynamodbav:"value"`
}


