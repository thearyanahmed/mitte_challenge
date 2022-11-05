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
	ID string `json:"id" bson:"_id"`
	Name string `json:"name"  bson:"name"`
}


