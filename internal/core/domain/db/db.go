package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// Store manages the set of API's for domain access.
type Store struct {
	log        *zap.SugaredLogger
	collection *mongo.Collection
}

// NewStore constructs a data for api access.
func NewStore(log *zap.SugaredLogger, collection *mongo.Collection) Store {
	return Store{
		log:        log,
		collection: collection,
	}
}

//TODO: add additional functionality as necessary

func (s Store) Update(ctx context.Context, filter bson.M, update bson.D, options options.UpdateOptions) (interface{}, error) {
	result, err := s.collection.UpdateOne(ctx, filter, update, &options)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s Store) Get(ctx context.Context, filter bson.M) (bson.M, error) {
	item := bson.M{}
	err := s.collection.FindOne(ctx, filter).Decode(&item)
	if err != nil {
		return nil, err
	}
	return item, nil
}
