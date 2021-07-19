package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kyh0703/golang/mongo/mongo"
	"go.mongodb.org/mongo-driver/mongo/bson"
)

func create() {
	conn := mongo.GetClient()
	database := conn.Database("testkyh")
	collection := database.Collection("podcats")

	res, err := collection.InsertOne(context.TODO(), bson.D{
		{Key: "test", Value: "1"},
		{Key: "test1", Value: "2"},
		{Key: "test2", Value: "3"},
		{Key: "test3", Value: "4"},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("INSERT RESULT", res)
}

func remove() {
	conn := mongo.GetClient()
	database := conn.Database("testkyh")
	collection := database.Collection("podcats")

	insertResult, err := writeDB.InsertOne(context.TODO(), bson.D{
		{Key: "logType", Value: "hihihi"},
		{Key: "logValue", Value: "hoihi"},
		{Key: "dbPrefix", Value: "hhihihi"},
		{Key: "time", Value: "dada"},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INSERT RESULT", insertResult)
}

func update() {
	conn := mongo.GetClient()
	database := conn.Database("testkyh")
	collection := database.Collection("podcats")

	res, err := writeDB.InsertOne(context.TODO(), bson.D{
		{Key: "test", Value: "1"},
		{Key: "test1", Value: "2"},
		{Key: "test2", Value: "3"},
		{Key: "test3", Value: "4"},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INSERT RESULT", res)
}

func delete() {
	conn := mongo.GetClient()
	database := conn.Database("testkyh")
	collection := database.Collection("podcats")

	insertResult, err := writeDB.InsertOne(context.TODO(), bson.D{
		{Key: "logType", Value: "hihihi"},
		{Key: "logValue", Value: "hoihi"},
		{Key: "dbPrefix", Value: "hhihihi"},
		{Key: "time", Value: "dada"},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INSERT RESULT", insertResult)
}

func main() {
	conn := mymongo.Connect("localhost", "27017")

	database := conn.Database("testkyh")
	collection := database.Collection("podcats")
	insert(collection)
}
