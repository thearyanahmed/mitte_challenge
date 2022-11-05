package repository

import (
	"context"
	"github.com/thearyanahmed/mitte_challenge/pkg/schema"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
)

const swipesTable = "swipes"

type SwipeRepository struct {
	db *dynamodb.Client
}

func NewSwipeRepository(db *dynamodb.Client) *SwipeRepository {
	return &SwipeRepository{
		db: db,
	}
}

func (r *SwipeRepository) InsertSwipe(ctx context.Context, schemaSchema schema.SwipeSchema) (schema.SwipeSchema, error) {
	// @todo handle in a central place
	// a schemaSchema could be extended to have these methods
	
	schemaSchema.ID = uuid.New().String()
	schemaSchema.CreatedAt = time.Now()
	attribute, err := attributevalue.MarshalMap(schemaSchema)

	if err != nil {
		return schema.SwipeSchema{}, err
	}

	_, err = r.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(swipesTable),
		Item:      attribute,
	})

	return schemaSchema, err
}

func (r SwipeRepository) GetSwipesByUserId(ctx context.Context, userId string) ([]schema.SwipeSchema, error) {
	filt := expression.Name("swiped_by").Equal(expression.Value(userId))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	if err != nil {
		// @todo these should to send raw error
		return []schema.SwipeSchema{}, err
	}

	result, err := r.db.Scan(ctx, &dynamodb.ScanInput{
		TableName:                 aws.String(swipesTable),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})

	if err != nil {
		return []schema.SwipeSchema{}, err
	}

	var collection []schema.SwipeSchema

	err = attributevalue.UnmarshalListOfMaps(result.Items, &collection)

	return collection, err
}

func (r SwipeRepository) CheckIfSwipeExists(ctx context.Context, userId, profileOwnerId string) (schema.SwipeSchema, bool, error) {

	expr, err := expression.NewBuilder().WithFilter(
		expression.And(
			expression.Name("swiped_by").Equal(expression.Value(userId)),
			expression.Name("profile_owner_id").Equal(expression.Value(profileOwnerId))),
		).Build()

	if err != nil {
		return schema.SwipeSchema{}, false, nil
	}

	result, err := r.db.Scan(ctx, &dynamodb.ScanInput{
		TableName:                 aws.String(swipesTable),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})

	if err != nil {
		return schema.SwipeSchema{}, false, nil
	}

	if result.Count < 1 {
		return schema.SwipeSchema{}, false, nil
	}

	swipe := schema.SwipeSchema{}

	for _, v := range result.Items {
		err = attributevalue.UnmarshalMap(v, &swipe)
		break
	}

	if err != nil {
		return schema.SwipeSchema{}, false, nil
	}

	return swipe, true, nil
}