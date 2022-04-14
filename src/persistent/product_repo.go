package persistent

import (
	"context"
	"mongo-store/src/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

var repo *mongo.Collection

func InitRepository(db *mongo.Database) {
	repo = db.Collection("users")
}

func Insert(p domain.Product) *mongo.InsertOneResult {
	result, err := repo.InsertOne(context.TODO(), p)
	if err != nil {
		panic(err)
	}
	return result
}
