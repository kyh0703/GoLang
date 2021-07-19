package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type logEntity struct {
	logType  string `bson:"logType"`
	logValue string `bson:"logValue"`
	dbPrefix string `bson:"dbPrefix"`
	time     string `bson:"time"`
}

func mongoConn() (client *mongo.Client) {
	// credential := options.Credential{
	// 	Username: "<USER_NAME>",
	// 	Password: "<PASSWORD>",
	// }

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Made")
	return client
}

func logWrite(writeDB *mongo.Collection) {
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
	conn := mongoConn()
	mongo := conn.Database("log_data").Collection("logs")

	fmt.Println(conn, mongo)
}
