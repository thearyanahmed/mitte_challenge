package repository

import (
	"context"
	"github.com/thearyanahmed/mitte_challenge/pkg/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const SwipesCollection = "swipes"

type SwipeRepository struct {
	collection *mongo.Collection
}

func NewSwipeRepository(db *mongo.Collection) *SwipeRepository {
	return &SwipeRepository{
		collection: db,
	}
}

func (r *SwipeRepository) InsertSwipe(ctx context.Context, swipe schema.SwipeSchema) (schema.SwipeSchema, error) {
	swipe.CreatedAt = time.Now()

	result , err := r.collection.InsertOne(ctx, swipe)

	if err != nil {
		return schema.SwipeSchema{}, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return schema.SwipeSchema{}, err
	}

	swipe.ID = oid

	return swipe, err
}

func (r SwipeRepository) GetSwipesByUserId(ctx context.Context, userId string) ([]schema.SwipeSchema, error) {
	filter := bson.D{{"swiped_by", userId}}

	cursor, err := r.collection.Find(ctx, filter)

	if err != nil {
		return []schema.SwipeSchema{}, err
	}

	var results []schema.SwipeSchema

	if err = cursor.All(ctx, &results); err != nil {
		return []schema.SwipeSchema{}, err
	}

	return results, nil
}

func (r SwipeRepository) CheckIfSwipeExists(ctx context.Context, swipedById, profileOwnerId string) (schema.SwipeSchema, bool, error) {
	filters := bson.D{
		{"$and",
			bson.A{
				bson.D{{"swiped_by", swipedById}},
				bson.D{{"profile_owner_id", profileOwnerId}},
			}},
	}

	queryFilter := bson.D{{"$and", bson.A{filters}}}

	var result schema.SwipeSchema
	err := r.collection.FindOne(ctx, queryFilter).Decode(&result)

	if err != nil {
		return schema.SwipeSchema{}, false, err
	}

	return result, true, nil
}