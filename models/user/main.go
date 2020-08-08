package user

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	collection *mongo.Collection
}

func (users *Users) Insert(ctx context.Context, user *User) error {
	_, err := users.collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Printf("Error while inserting user: %s", err.Error())
		return err
	}
	return nil
}

func (users *Users) Get(ctx context.Context, username, password string) (*User, error) {
	var foundUser User
	filter := bson.M{"username": username, "password": password}
	err := users.collection.FindOne(ctx, filter).Decode(&foundUser)
	if err != nil {
		fmt.Printf("Error while getting user: %s", err.Error())
		return nil, err
	}
	return &foundUser, nil
}

func NewUsers(db *mongo.Database, collName string) *Users {
	return &Users{collection: db.Collection(collName)}
}