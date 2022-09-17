package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	uriString := os.Getenv("URI_STRING")

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uriString).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	dbs, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(dbs)

	//friends := client.Database("friends")
	//males := friends.Collection("males")
	//females := friends.Collection("females")
	todo := client.Database("todo")
	myTodos := todo.Collection("todos")
	urgentTasks := todo.Collection("Task")
	todayJobs, err := urgentTasks.InsertMany(ctx, []interface{}{
		bson.D{{"task", "wash body"}},
		bson.D{{"task", "wash car"}},
		bson.D{{"task", "wash floors"}},
		bson.D{{"task", "wash dog"}},
		bson.D{{"task", "make dinner"}},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(todayJobs.InsertedIDs)

	firstTodo, err := myTodos.InsertOne(ctx, bson.D{
		{"task", "get off your arse!"},
	})
	if err != nil {
		log.Fatal(err)
	}

	//firstMale, err := males.InsertOne(ctx, bson.D{
	//	{"name", "Drake"},
	//	{"age", 32},
	//	{"worth", 19726118},
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//firstFemale, err := females.InsertOne(ctx, bson.D{
	//	{"name", "Doja"},
	//	{"age", 29},
	//	{"worth", 8226194},
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(firstMale.InsertedID)
	//fmt.Println(firstFemale.InsertedID)
	fmt.Println(firstTodo.InsertedID)
}
