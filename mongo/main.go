package main

import (
	"context"
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/kyh0703/golang/mongo/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type Config struct {
	MongoServer string `env:"MONGO_SERVER" envDefault:"127.0.0.1"`
	MongoPort   string `env:"MONGO_PORT" envDefault:"27017"`
	MongoId     string `env:"MONGO_ID" envDefault:"admin"`
	MongoPwd    string `env:"MONGO_PWD" envDefault:"dnflth"`
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

	fmt.Println("insert result", res)
}

func createCondition(googleID, name string) {
	conn := mongo.GetClient()
	database := conn.Database("testkyh")
	collection := database.Collection("podcats")

	filter := bson.M{"googleid": googleID, "name": name}
	num, _ := collection.CountDocuments(context.TODO(), filter)
	if num != 0 {
		return
	}

	type insertData struct {
		GoogleId string
		Name     string
	}

	newData := insertData{
		GoogleId: googleID,
		Name:     name,
	}

	res, err := collection.InsertOne(context.TODO(), newData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[insert result]\n", res)
}

func read() {
	conn := mongo.GetClient()
	database := conn.Database("testkyh")
	collection := database.Collection("podcats")

	filter := bson.M{}
	res, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var data []bson.M
	if err := res.All(context.TODO(), &data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("[read result]\n", data)
}

func readCondition(googleID string) {
	conn := mongo.GetClient()
	database := conn.Database("testkyh")
	collection := database.Collection("podcats")

	filter := bson.M{"googleid": googleID}
	res, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var data []bson.M
	if err := res.All(context.TODO(), &data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("[read condition result]\n", data)
}

func update(googleID, updateName string) {
	conn := mongo.GetClient()
	database := conn.Database("testkyh")
	collection := database.Collection("podcats")

	filter := bson.M{"googleid": googleID}

	// update : $set
	// add list: $push
	// del list: $pull
	update := bson.M{
		"$set": bson.M{
			"name": updateName,
		},
	}

	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[update result]\n", res)
}

func delete(googleID string) {
	conn := mongo.GetClient()
	database := conn.Database("testkyh")
	collection := database.Collection("podcats")

	filter := bson.M{"googleid": googleID}
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[delete result]\n", res)
}

func main() {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%v", err)
	}

	if err := mongo.Connect(cfg.MongoServer, cfg.MongoPort); err != nil {
		log.Fatal("Connect MongoDB Fail")
	}

	fmt.Println("connected mongo")
	fmt.Println("----------------> create")
	create()
	fmt.Println("----------------> create condition")
	createCondition("test", "1234")
	fmt.Println("----------------> read all")
	read()
	fmt.Println("----------------> read all by id")
	readCondition("test")
	fmt.Println("----------------> update")
	update("test", "kimyeonho")
	fmt.Println("----------------> delete")
	delete("test1")
}
