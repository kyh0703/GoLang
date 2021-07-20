package main

import (
	"context"
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/kyh0703/golang/mongo/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type Config struct {
	MongoHost string `env:"MONGO_INIT_SERVER" envDefault: "127.0.0.1"`
	MongoPort string `env:"MONGO_INIT_PORT" envDefault: "27017"`
	MongoId   string `env:"MONGO_INIT_ID" envDefault: "admin"`
	MongoPwd  string `env:"MONGO_INIT_PWD" envDefault: "dnflth"`
}

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

func read() {
	// conn := mongo.GetClient()
	// database := conn.Database("testkyh")
	// collection := database.Collection("podcats")
}

func update() {
	// conn := mongo.GetClient()
	// database := conn.Database("testkyh")
	// collection := database.Collection("podcats")
}

func delete() {
	// conn := mongo.GetClient()
	// database := conn.Database("testkyh")
	// collection := database.Collection("podcats")
}

func main() {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println(cfg.MongoHost)
	if err := mongo.Connect(cfg.MongoHost, cfg.MongoPort); err != nil {
		log.Fatal("Connect MongoDB Fail")
	}
	// time.Sleep(time.Second * 10)
	create()
}
