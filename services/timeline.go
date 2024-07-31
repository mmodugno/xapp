package services

type Timeline struct {
	UserID string  `json:"user_id"`
	Tweets []Tweet `json:"tweets"`
}

func (tm *Timeline) GetTimeline(id string) (Timeline, error) {
	var user User
	var tweet Tweet

	following, err := user.getFollowing(id)
	if err != nil {
		return Timeline{}, err
	}

	tweets, err := tweet.GetTweetsOfUser(following[0]) //TODO validate all followings
	if err != nil {
		return Timeline{}, err
	}
	return Timeline{
		UserID: id,
		Tweets: tweets,
	}, nil
}
