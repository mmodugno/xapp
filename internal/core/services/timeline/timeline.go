package timeline

import (
	"sort"
	"x-app-go/internal/core/services/tweets"
	"x-app-go/internal/core/services/users"
)

type Timeline struct {
	User  users.Client
	Tweet tweets.Client
}

type TimelineOfUser struct {
	UserID string         `json:"user"`
	Tweets []tweets.Tweet `json:"tweets"`
}

func New() *Timeline {
	return &Timeline{
		User:  &users.User{},
		Tweet: &tweets.Tweet{},
	}
}

func (tm *Timeline) GetTimeline(id string) (*TimelineOfUser, error) {
	following, err := tm.User.GetFollowing(id)
	if err != nil {
		return nil, err
	}
	var tweets []tweets.Tweet

	for _, value := range following {
		ts, err := tm.Tweet.GetTweetsOfUser(value)
		if err != nil {
			return nil, err
		}
		tweets = append(tweets, ts...)
		println(tweets)
	}
	//sort tweets by created at
	sort.Slice(tweets, func(i, j int) bool {
		return tweets[i].CreatedAt.After(tweets[j].CreatedAt)
	})

	return &TimelineOfUser{
		UserID: id,
		Tweets: tweets,
	}, nil
}
