package users

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"regexp"
	"slices"
	"time"
	"x-app-go/internal/core/services"
)

type User struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string    `json:"username" bson:"username"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Following []string  `json:"following,omitempty" bson:"following,omitempty"`
}

type Client interface {
	InsertUser(username string) (string, error)
	GetUserByID(id string) (User, error)
	GetUserByUsername(username string) (User, error)
	Delete(id string) error
	FollowUser(id string, username string) error
	GetFollowing(id string) ([]string, error)
}

func (u *User) InsertUser(username string) (string, error) {
	c := services.CollectionPointer("users")

	now := time.Now()

	err := checkUsername(username)
	if err != nil {
		return "", err
	}

	_, err = u.GetUserByUsername(username)
	if err == nil {
		return "", fmt.Errorf("username is already used")
	}

	result, err := c.InsertOne(context.TODO(), User{
		Username:  username,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		fmt.Println("error when creating user: ", err)
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID)
	idString := id.Hex()

	return idString, nil
}

func (u *User) GetUserByID(id string) (User, error) {
	c := services.CollectionPointer("users")
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
	c := services.CollectionPointer("users")

	var user User
	err := c.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *User) Delete(id string) error {
	c := services.CollectionPointer("users")

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = c.DeleteOne(context.Background(), bson.M{"_id": mongoID})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *User) FollowUser(id string, username string) error {
	c := services.CollectionPointer("users")

	//Get user information
	user, err := u.GetUserByID(id)
	if err != nil {
		return fmt.Errorf("cannot get logged user: %s", err.Error())
	}
	f := user.Following

	//get ID of username received
	follow, err := u.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("cannot get user: %s", err.Error())
	}

	//checks if it already follows
	if slices.Contains(f, follow.ID) {
		return fmt.Errorf("already following user")
	}

	//add following username to following var
	f = append(f, follow.ID)
	mongoID, err := primitive.ObjectIDFromHex(id)

	update := bson.D{
		{"$set", bson.D{
			{"following", f},
		}},
	}

	_, err = c.UpdateOne(
		context.Background(),
		bson.M{"_id": mongoID},
		update,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *User) GetFollowing(id string) ([]string, error) {
	log.Print("get following")
	user, err := u.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	log.Print(user.Following)
	return user.Following, nil
}

func checkUsername(username string) error {
	usernameConvention := `^[a-z_]([a-z0-9_-]{0,20}|[a-z0-9_-]{0,20}\$)$`

	re := regexp.MustCompile(usernameConvention).MatchString
	if !re(username) {
		return fmt.Errorf("invalid username")
	}
	return nil
}
