package domains

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateDelivered(collection *mongo.Collection, name string) error {

	filter := bson.M{
		"name": name,
	}
	update := bson.D{
		{"$inc", bson.D{{"delivered", 1}}},
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := collection.FindOneAndUpdate(context.TODO(), filter, update, &opt)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
func UpdateBounced(collection *mongo.Collection, name string) error {

	filter := bson.M{
		"name": name,
	}
	update := bson.D{
		{"$inc", bson.D{{"bounced", 1}}},
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := collection.FindOneAndUpdate(context.TODO(), filter, update, &opt)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func CheckStatus(collection *mongo.Collection, name string) (string, error) {
	domain := bson.M{}
	err := collection.FindOne(context.TODO(), bson.M{
		"name": name,
	}).Decode(&domain)
	if err != nil {
		return "", err
	}
	bounced := domain["bounced"]
	delivered := domain["delivered"]
	if bounced == nil {
		bounced = 0
	}
	if delivered == nil {
		delivered = 0
	}
	bouncedInt, ok := bounced.(int32)
	if !ok {
		error := errors.New("Bounced did not convert to integer")
		return "", error
	}
	deliveredInt, ok := delivered.(int32)
	if !ok {
		error := errors.New("Delivered did not convert to integer")
		return "", error
	}
	if bouncedInt >= 1 {
		return "not a catch-all", nil
	}
	if deliveredInt >= 1000 {
		return "catch-all", nil
	} else {
		return "unknown", nil
	}
}
