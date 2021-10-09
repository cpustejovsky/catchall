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
	log.Println(name)
	_, err := d.DB.UpdateOne(context.TODO(), bson.M{
		"name": name,
	}, bson.D{
		{"$inc", bson.D{{"delivered", 1}}},
	}, options.Update().SetUpsert(true))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (d *DomainModel) UpdateBounced(name string) error {
	log.Println(name)
	_, err := d.DB.UpdateOne(context.TODO(), bson.M{
		"name": name,
	}, bson.D{
		{"$inc", bson.D{{"bounced", 1}}},
	}, options.Update().SetUpsert(true))
	if err != nil {
		log.Println(err)
		return err
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
