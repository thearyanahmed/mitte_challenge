package repository

import (
	"context"
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

func (r *SwipeRepository) InsertSwipe(ctx context.Context, schema SwipeSchmea) (SwipeSchmea, error) {
	// @todo handle in a central place
	// a schema could be extended to have these methods
	
	schema.ID = uuid.New().String()
	schema.CreatedAt = time.Now()
	attribute, err := attributevalue.MarshalMap(schema)

	if err != nil {
		return SwipeSchmea{}, err
	}

	_, err = r.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(swipesTable),
		Item:      attribute,
	})

	return schema, err
}

func (r SwipeRepository) GetSwipesByUserId(ctx context.Context, userId string) ([]SwipeSchmea, error) {
	filt := expression.Name("swiped_by").Equal(expression.Value(userId))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	if err != nil {
		// @todo these should to send raw error
		return []SwipeSchmea{}, err
	}

	result, err := r.db.Scan(ctx, &dynamodb.ScanInput{
		TableName:                 aws.String(swipesTable),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})

	if err != nil {
		return []SwipeSchmea{}, err
	}

	var collection []SwipeSchmea

	err = attributevalue.UnmarshalListOfMaps(result.Items, &collection)

	return collection, err
}

func (r SwipeRepository) CheckIfSwipeExists(ctx context.Context, userId, profileOwnerId string) (SwipeSchmea, bool, error) {

	expr, err := expression.NewBuilder().WithFilter(
		expression.And(
			expression.Name("swiped_by").Equal(expression.Value(userId)),
			expression.Name("profile_owner_id").Equal(expression.Value(profileOwnerId))),
		).Build()

	if err != nil {
		return SwipeSchmea{}, false, nil
	}

	result, err := r.db.Scan(ctx, &dynamodb.ScanInput{
		TableName:                 aws.String(swipesTable),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})

	if err != nil {
		return SwipeSchmea{}, false, nil
	}

	if result.Count < 1 {
		return SwipeSchmea{}, false, nil
	}

	swipe := SwipeSchmea{}

	for _, v := range result.Items {
		err = attributevalue.UnmarshalMap(v, &swipe)
		break
	}

	if err != nil {
		return SwipeSchmea{}, false, nil
	}

	return swipe, true, nil
}