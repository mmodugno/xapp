package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type User struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string    `json:"username" bson:"username"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

var client *mongo.Client

func New(mongo *mongo.Client) User {
	client = mongo
	return User{}
}

func CollectionPointer(collection string) *mongo.Collection {
	return client.Database("xapp_db").Collection(collection)
}

func (u *User) InsertUser(entry User) error {
	c := CollectionPointer("users")

	now := time.Now()

	_, err := u.GetUserByUsername(entry.Username)
	if err == nil {
		return fmt.Errorf("username is already used")
	}

	_, err = c.InsertOne(context.TODO(), User{
		Username:  entry.Username,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		fmt.Println("error when creating user: ", err)
		return err
	}
	return nil
}

func (u *User) GetUserByID(id string) (User, error) {
	c := CollectionPointer("users")
	var user User

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return User{}, err
	}

	err = c.FindOne(context.Background(), bson.M{"_id": mongoID}).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *User) GetUserByUsername(username string) (User, error) {
	c := CollectionPointer("users")

	var user User
	err := c.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
