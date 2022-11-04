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

const swipes_table = "swipes"

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
		TableName: aws.String(swipes_table),
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
		TableName:                 aws.String(swipes_table),
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
