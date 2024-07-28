package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"x-app-go/db"
	"x-app-go/handlers"
	"x-app-go/services"
)

type Application struct {
	Models services.Todo
}

func main() {
	fmt.Println("it works")

	mongoClient, err := db.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}

	//we need the context to execute querys, we need cancel to stop them
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Connect(ctx); err != nil {
			panic(err)
		}
	}()

	services.New(mongoClient)
	log.Println("server running in port ", 8080)
	log.Fatal(http.ListenAndServe(":8080", handlers.CreateRouter()))
}
