package mongodb

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type DomainModel struct {
	DB *mongo.Collection
}

func (d *DomainModel) UpdateDelivered(name string) error {
	log.Println(name)
	return nil
}
func (d *DomainModel) UpdateBounced(name string) error {
	return nil
}
func (d *DomainModel) CheckStatus(name string) (string, error) {
	return "foo", nil
}
