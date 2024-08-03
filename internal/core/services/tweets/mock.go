package tweets

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mock struct {
	mock.Mock
}

//func (t *Mock) GetTweetsOfUser(userID string) ([]Tweet, error) {
//	args := t.Called(userID)
//	if err := args.Get(1); err != nil {
//		var res []Tweet
//		return res, err.(error)
//	}
//	return args.Get(0).([]Tweet), nil
//}
//
//func (t *Mock) InsertTweet(id string, entry Tweet) error {
//	args := t.Called(id, entry)
//	if err := args.Get(0); err != nil {
//		return err.(error)
//	}
//	return nil
//}

type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}
