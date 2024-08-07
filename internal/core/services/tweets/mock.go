package tweets

import (
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (t *Mock) GetTweetsOfUser(userID string) ([]Tweet, error) {
	args := t.Called(userID)
	if err := args.Get(1); err != nil {
		var res []Tweet
		return res, err.(error)
	}
	return args.Get(0).([]Tweet), nil
}

func (t *Mock) InsertTweet(id string, entry Tweet) error {
	args := t.Called(id, entry)
	if err := args.Get(0); err != nil {
		return err.(error)
	}
	return nil
}
