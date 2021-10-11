package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DomainModelV2 struct {
	DB *mongo.Database
}

func (d *DomainModelV2) UpdateDelivered(name string) error {
	deliveredDomains := d.DB.Collection("delivered_domains")
	item := bson.D{{Key: "domain_name", Value: name}}

	_, err := deliveredDomains.InsertOne(context.TODO(), item)
	if err != nil {
		return err
	}
	return nil
}
func (d *DomainModelV2) UpdateBounced(name string) error {
	bouncedDomains := d.DB.Collection("bounced_domains")
	item := bson.D{{Key: "domain_name", Value: name}}
	_, err := bouncedDomains.InsertOne(context.TODO(), item)
	if err != nil {
		return err
	}
	return nil
	return nil
}

func (d *DomainModelV2) CheckStatus(name string) (string, error) {
	// deliveredDomains := d.DB.Collection("delivered_domains")
	// bouncedDomains := d.DB.Collection("bounced_domains")
	// matchStage := bson.D{{"$match", bson.D{{"domain_name", name}}}}
	// groupStage := bson.D{{"$group", bson.D{{"name", "$domain_name"}, {"total", bson.D{{"$counbt", "$duration"}}}}}}

	// domain := bson.M{}

	// err := d.DB.FindOne(context.TODO(), bson.M{
	// 	"name": name,
	// }).Decode(&domain)
	// if err != nil {
	// 	log.Println(err)
	// 	return "", err
	// }
	// bounced := domain["bounced"]
	// delivered := domain["delivered"]
	// if bounced == nil {
	// 	bounced = 0
	// }
	// if delivered == nil {
	// 	delivered = 0
	// }
	// fmt.Printf("bounced = %T\n", bounced)
	// fmt.Printf("delivered = %T\n", delivered)
	// bouncedInt, ok := bounced.(int32)
	// if !ok {
	// 	error := errors.New("Bounced did not convert to integer")
	// 	log.Println(error)
	// 	return "", error
	// }
	// deliveredInt, ok := delivered.(int32)
	// if !ok {
	// 	error := errors.New("Delivered did not convert to integer")
	// 	log.Println(error)
	// 	return "", error
	// }
	// if bouncedInt >= 1 {
	// 	return "not a catch-all", nil
	// }
	// if deliveredInt >= 1000 {
	// 	return "catch-all", nil
	// } else {
	// 	return "unknown", nil
	// }
	return "foo", nil
}
