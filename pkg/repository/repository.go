package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func newObjectId() primitive.ObjectID {
	return primitive.NewObjectID()
}
