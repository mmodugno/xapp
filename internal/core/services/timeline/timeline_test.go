package timeline

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"x-app-go/internal/core/services/tweets"
	"x-app-go/internal/core/services/users"
)

func Test_Timeline_GetTimeline(t *testing.T) {
	type args struct {
		id string
	}
	time := time.Now()

	tests := []struct {
		name       string
		args       args
		usersMock  users.Mock
		tweetsMock tweets.Mock
		usersInit  func(in *users.Mock)
		tweetsInit func(in *tweets.Mock)
		want       assert.ValueAssertionFunc
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name: "given an user with 2 following should return those tweets",
			args: args{
				id: "123",
			},
			usersInit: func(in *users.Mock) {
				in.On("GetFollowing", "123").Return([]string{"1", "2"}, nil)
			},
			tweetsInit: func(in *tweets.Mock) {
				in.On("GetTweetsOfUser", "1").Return([]tweets.Tweet{
					{ID: "tweet 1 usuario 1"},
					{ID: "tweet 2 del usuario 1"},
				}, nil,
				)
				in.On("GetTweetsOfUser", "2").Return([]tweets.Tweet{
					{ID: "tweet 1 usuario 2"},
				}, nil,
				)
			},
			want: func(t assert.TestingT, got interface{}, i2 ...interface{}) bool {
				want := &TimelineOfUser{
					UserID: "123",
					Tweets: []tweets.Tweet{
						{ID: "tweet 1 usuario 1"},
						{ID: "tweet 2 del usuario 1"},
						{ID: "tweet 1 usuario 2"},
					},
				}
				return assert.Equal(t, want, got)
			},
			wantErr: assert.NoError,
		},
		{
			name: "given an user with 2 following should return those tweets in recent order of createdDate",
			args: args{
				id: "123",
			},
			usersInit: func(in *users.Mock) {
				in.On("GetFollowing", "123").Return([]string{"1", "2"}, nil)
			},
			tweetsInit: func(in *tweets.Mock) {
				in.On("GetTweetsOfUser", "1").Return([]tweets.Tweet{
					{ID: "tweet 1 usuario 1", CreatedAt: time.AddDate(0, 0, 1)},
					{ID: "tweet 2 del usuario 1", CreatedAt: time},
				}, nil,
				)
				in.On("GetTweetsOfUser", "2").Return([]tweets.Tweet{
					{ID: "tweet 1 usuario 2", CreatedAt: time.AddDate(0, 0, -1)},
				}, nil,
				)
			},
			want: func(t assert.TestingT, got interface{}, i2 ...interface{}) bool {
				want := &TimelineOfUser{
					UserID: "123",
					Tweets: []tweets.Tweet{
						{ID: "tweet 1 usuario 1", CreatedAt: time.AddDate(0, 0, 1)},
						{ID: "tweet 2 del usuario 1", CreatedAt: time},
						{ID: "tweet 1 usuario 2", CreatedAt: time.AddDate(0, 0, -1)},
					},
				}
				return assert.EqualValues(t, want, got)
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.tweetsInit(&tt.tweetsMock)
			tt.usersInit(&tt.usersMock)

			tm := &Timeline{
				User:  &tt.usersMock,
				Tweet: &tt.tweetsMock,
			}
			got, err := tm.GetTimeline(tt.args.id)
			tt.want(t, got)
			tt.wantErr(t, err)
		})
	}
}
