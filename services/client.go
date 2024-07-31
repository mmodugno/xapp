package services

import "go.mongodb.org/mongo-driver/mongo"

var UserClient *mongo.Client

func CollectionPointer(collection string) *mongo.Collection {
	return UserClient.Database("xapp_db").Collection(collection)
}

func New(mongo *mongo.Client) {
	UserClient = mongo
	return
}
