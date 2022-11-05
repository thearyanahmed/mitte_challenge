package schema


var (
	Traits []TraitsSchema
)

func init() {
	Traits = append(Traits, TraitsSchema{
		ID:    "",
		Name: "",
	})
}

type TraitsSchema struct {
	ID string `json:"id" dynamodbav:"id"`
	Name string `json:"name"  dynamodbav:"name"`
}


