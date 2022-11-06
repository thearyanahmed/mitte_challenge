package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/thearyanahmed/mitte_challenge/pkg/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UsersCollection = "users"

// UserRepository represents the user repository that communicates with the database.
type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: db,
	}
}

func (r *UserRepository) Insert(ctx context.Context, user *schema.UserSchema) (string, error) {
	user.ID = newObjectId()
	user.CreatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return "", errors.New("could not convert inserted id to primitive id")
	}

	return oid.Hex(), nil
}

func (r *UserRepository) FindById(ctx context.Context, hex string) (schema.UserSchema, error) {
	objectId, err := primitive.ObjectIDFromHex(hex)

	if err != nil {
		return schema.UserSchema{}, err
	}

	var user schema.UserSchema

	if err := r.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user); err != nil {
		return schema.UserSchema{}, nil
	}

	return user, nil
}

// FindByEmail @todo change returns, manage errors better
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (schema.UserSchema, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var user schema.UserSchema

	if err := r.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return schema.UserSchema{}, nil
	}

	return user, nil
}

func (r *UserRepository) Find(ctx context.Context, requestFilters map[string]string) ([]schema.UserSchema, error) {
	filters := mapPropertyFilter(requestFilters)

	cursor, err := r.collection.Find(ctx, filters)

	if err != nil {
		return []schema.UserSchema{}, err
	}

	var results []schema.UserSchema

	if err = cursor.All(ctx, &results); err != nil {
		return []schema.UserSchema{}, err
	}

	return results, nil
}

func (r *UserRepository) FindMatch() {
	/**
	db.users.aggregate([
	  {
	    '$match' : {
	      '$and': [
	        { 'gender': 'male' },
	        { 'age' : { '$lte': 100 } },
	        {
	          '$or' : [
	              {'traits.id' : '1'},
	              {'traits.id' : '4'}
	            ]
	          }
	      ]
	    },
	  },
	  {
	    '$project' :  {
	      attractiveness_score : { '$sum' : '$traits.value' },
	      name: '$name',
	      email: '$email',
	      age: '$age',
	      gender: '$gender',
	      traits: '$traits'
	    }
	  }
	]).sort({ 'attractiveness_score' : -1 })
	*/

	project := bson.D{
		{"name", "$name"},
		{"email", "$email"},
		{"age", "$age"},
		{"gender", "$gender"},
		{"traits", "$traits"},
		{"attractiveness_score", bson.D{
			{"$sum", "$traits.value"},
		}},
	}

	sort := bson.D{
		{"attractiveness_score", -1},
	}

	match := bson.D{{
		"$and", bson.A{
			bson.D{
				{"gender", "male"},
				{"age", bson.D{{"$lte", 100}}},
				{"$or", bson.A{
					bson.D{{"traits.id", "1"}},
					bson.D{{"traits.id", "4"}},
				}},
			},
		},
	}}

	_ = bson.D{
		{"$and", bson.D{
			{"gender", "male"},
			{"age", bson.D{
				{"$lte", 100},
			}},
			{"$or", bson.A{
				bson.D{{"traits.id", "1"}},
				bson.D{{"traits.id", "4"}},
			}},
		}},
	}

	//
	//mongo.Pipeline{
	//	//		{{"$group", bson.D{{"_id", "$state"}, {"totalPop", bson.D{{"$sum", "$pop"}}}}}},
	//	//		{{"$match", bson.D{{"totalPop", bson.D{{"$gte", 10*1000*1000}}}}}},
	//	//	}

	pipeline := mongo.Pipeline{
		{{"$match", match}},
		{{"$project", project}},
		{{"$sort", sort}},
	}

	cursor, err := r.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		fmt.Println("EERRR_.1", err)
		return
	}

	var results []schema.UserSchema
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println("EERRR_.2", err)
		return
	}

	fmt.Println("HERE", len(results))
	for _, user := range results {
		fmt.Println("USER", user.ID, user.Name, user.Email, user.Gender, user.Traits)
	}
}

func mapPropertyFilter(requestFilters map[string]string) bson.M {
	x := bson.M{}

	for k, v := range requestFilters {
		if numeric, err := strconv.Atoi(v); err == nil {
			x[k] = numeric
		} else {
			x[k] = v
		}
	}

	return x
}
