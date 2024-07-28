package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Todo struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string    `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool      `json:"completed" bson:"completed"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

var client *mongo.Client

func New(mongo *mongo.Client) Todo {
	client = mongo

	return Todo{}
}

func returnCollectionPointer(collection string) *mongo.Collection {
	return client.Database("todos_db").Collection(collection)
}

func (t *Todo) InsertTodo(entry Todo) error {
	c := returnCollectionPointer("todos")

	now := time.Now()
	_, err := c.InsertOne(context.TODO(), Todo{
		Task:      entry.Task,
		Completed: entry.Completed,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	return nil
}

func (t *Todo) GetAllTodos() ([]Todo, error) {
	c := returnCollectionPointer("todos")
	var todos []Todo

	cursor, err := c.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal("Error: ", err)
		return nil, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.Background()) {
		var todo Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *Todo) GetTodoById(id string) (Todo, error) {
	c := returnCollectionPointer("todos")
	var todo Todo

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Todo{}, err
	}

	err = c.FindOne(context.Background(), bson.M{"_id": mongoID}).Decode(&todo)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

func (t *Todo) UpdateTodo(id string, entry Todo) (*mongo.UpdateResult, error) {
	c := returnCollectionPointer("todos")

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.D{
		{"$set", bson.D{
			{"task", entry.Task},
			{"completed", entry.Completed},
			{"updated_at", time.Now()},
		}},
	}

	res, err := c.UpdateOne(
		context.Background(),
		bson.M{"_id": mongoID},
		update,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (t *Todo) DeleteTodo(id string) error {
	c := returnCollectionPointer("todos")

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
