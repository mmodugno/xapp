package users

import "github.com/stretchr/testify/mock"

type Mock struct {
	mock.Mock
}

func (u *Mock) InsertUser(username string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u *Mock) GetUserByUsername(username string) (User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *Mock) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (u *Mock) FollowUser(id string, username string) error {
	//TODO implement me
	panic("implement me")
}

func (u *Mock) GetUserByID(id string) (User, error) {
	args := u.Called(id)
	if err := args.Get(1); err != nil {
		return User{}, err.(error)
	}
	return args.Get(0).(User), nil
}

func (u *Mock) GetFollowing(id string) ([]string, error) {
	args := u.Called(id)
	if err := args.Get(1); err != nil {
		return nil, err.(error)
	}
	return args.Get(0).([]string), nil
}
