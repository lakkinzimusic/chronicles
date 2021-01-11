package service

import (
	//"go.mongodb.org/mongo-driver/bson"
	"context"
	"github.com/lakkinzimusic/chronicles/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	Create(context.Context, *model.Chronicle) (interface{}, error)
	All(context.Context) ([]*model.Chronicle, error)
}

func NewStore(collection *mongo.Collection) Store{
	return &MongoRepository{
		collection: collection,
	}
}

type MongoRepository struct {
	collection *mongo.Collection
}

func (s *MongoRepository) Create(ctx context.Context, chronicle *model.Chronicle) (interface{}, error) {
	item, err := s.collection.InsertOne(ctx, chronicle)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *MongoRepository) All(ctx context.Context) ([]*model.Chronicle, error) {
	var result []*model.Chronicle
	cursor, err := s.collection.Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		_ = cursor.Decode(result)
	}
	return result, nil
}