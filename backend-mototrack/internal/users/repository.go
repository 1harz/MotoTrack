package users

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client, databaseName, collectionName string) *UserRepository {
	collection := client.Database(databaseName).Collection(collectionName)
	return &UserRepository{collection: collection}
}

func (repo *UserRepository) Insert(user *User) error {
	_, err := repo.collection.InsertOne(context.TODO(), user)
	return err
}