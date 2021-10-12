package mongodb

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DomainModel struct {
	DB *mongo.Collection
}

func (d *DomainModel) UpdateDelivered(name string) error {

	filter := bson.M{
		"name": name,
	}
	update := bson.D{
		{"$inc", bson.D{{"delivered", 1}}},
	}

	upsert := true
	opt := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
	}

	result := d.DB.FindOneAndUpdate(context.TODO(), filter, update, &opt)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
func (d *DomainModel) UpdateBounced(name string) error {

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

	result := d.DB.FindOneAndUpdate(context.TODO(), filter, update, &opt)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (d *DomainModel) CheckStatus(name string) (string, error) {
	domain := bson.M{}
	err := d.DB.FindOne(context.TODO(), bson.M{
		"name": name,
	}).Decode(&domain)
	if err != nil {
		log.Println(err)
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
	fmt.Printf("bounced = %T\n", bounced)
	fmt.Printf("delivered = %T\n", delivered)
	bouncedInt, ok := bounced.(int32)
	if !ok {
		error := errors.New("Bounced did not convert to integer")
		log.Println(error)
		return "", error
	}
	deliveredInt, ok := delivered.(int32)
	if !ok {
		error := errors.New("Delivered did not convert to integer")
		log.Println(error)
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
	return "foo", nil
}
