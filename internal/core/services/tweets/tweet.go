package tweets

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"x-app-go/internal/core/services"
)

type TweetModel struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Content   string    `json:"content"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type Tweet struct {
	Collection *mongo.Client
}

type Client interface {
	InsertTweet(id string, entry TweetModel) error
	GetTweetsOfUser(userID string) ([]TweetModel, error)
}

func (t *Tweet) InsertTweet(id string, entry TweetModel) error {
	c := services.CollectionPointer("tweets")

	now := time.Now()

	_, err := c.InsertOne(context.TODO(), TweetModel{
		Content:   entry.Content,
		UserID:    id,
		CreatedAt: now,
	})
	if err != nil {
		fmt.Println("error when posting Tweet: ", err)
		return err
	}
	return nil
}

func (t *Tweet) GetTweetsOfUser(userID string) ([]TweetModel, error) {
	c := t.Collection.Database("xapp_db").Collection("tweets")
	var tweets []TweetModel

	sort := bson.D{{"created_at", 1}}
	opts := options.Find().SetSort(sort)

	cursor, err := c.Find(context.TODO(), bson.M{"userid": userID}, opts)
	if err != nil {
		fmt.Println("error when searching tweets: ", err)
		return nil, err
	}

	if err = cursor.All(context.TODO(), &tweets); err != nil {
		panic(err)
	}
	for _, result := range tweets {
		_, err := bson.MarshalExtJSON(result, false, false)
		if err != nil {
			return nil, err
		}
	}
	return tweets, nil
}
