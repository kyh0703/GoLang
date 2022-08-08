package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/kyh0703/golang/mongo/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Config struct {
	MongoServer string `env:"MONGO_SERVER" envDefault:"127.0.0.1"`
	MongoPort   string `env:"MONGO_PORT" envDefault:"27017"`
	MongoId     string `env:"MONGO_ID" envDefault:"admin"`
	MongoPwd    string `env:"MONGO_PWD" envDefault:"dnflth"`
}

type Contact struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Email  string             `bson:"email,omitempty"`
	Tester Email              `bson:"test,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

type Email struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	SendDate time.Time          `bson:"senddate,omitempty"`
	Subject  string             `bson:"subject,omitempty"`
	Content  string             `bson:"content,omitempty"`
}

type Sequence struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	Emails    []interface{}        `bson:"emails,omitempty"`
	Receivers []primitive.ObjectID `bson:"receivers,omitempty"`
}

// type Account struct {
// 	ID         primitive.ObjectID `bson:"_id"            json:"_id,omitempty"`
// 	Name       string             `bson:"name"           json:"name,omitempty"`
// 	Alias      string             `bson:"alias"          json:"alias,omitempty"`
// 	Enable     bool               `bson:"enable"         json:"enable,omitempty"`
// 	Emails     []interface{}      `bson:"emails,omitempty" json:email,omitempty"`
// 	ExpireDate time.Time          `bson:"expireDate"     json:"expireDate,omitempty"`
// 	Timezone   string             `bson:"timezone"       json:"timezone,omitempty"`
// 	IsMaster   bool               `bson:"isMaster"       json:"isMaster,omitempty"`
// }

// type Email struct {
// 	ID       primitive.ObjectID `bson:"_id"            json:"_id,omitempty"`
// 	SendDate time.Time          `bson:"senddate,omitempty"`
// 	Subject  string             `bson:"subject,omitempty"`
// 	Content  string             `bson:"content,omitempty"`
// }

func create(ctx context.Context) {
	// conn := mongo.GetClient()
	// database := conn.Database("testkyh")
	// collection := database.Collection("podcats")

	// manyContacts := []interface{}{
	// 	Contact{
	// 		Name:  "John Doe",
	// 		Email: "john@doe.com",
	// 		Tags:  []string{"Lead", "Subscriber"},
	// 	},
	// Contact{
	// 	Name:   "Mia May",
	// 	Email:  "mm@example.com",
	// 	Tester: "test입니다",
	// },
	// 	Contact{
	// 		Name:  "Lia Shmia",
	// 		Email: "lia@shmia.com",
	// 		Tags:  []string{"Lead", "Customer"},
	// 	},
	// }
	// res, err := collection.InsertOne(ctx, Contact{
	// 	Name:   "Mia May",
	// 	Email:  "mm@example.com",
	// 	Tester: "test입니다",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("insert result", res)
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

type CallSegment struct {
	Event    string    `json:"event,omitempty"`       // 이벤트
	CallID   string    `json:"call-id,omitempty"`     // 콜 아이디
	Type     int       `json:"type"`                  // 콜 타입
	End      *End      `json:"end,omitempty"`         // 종료 정보
	CallTime *CallTime `json:"calltime,omitempty"`    // 콜 시간 정보
	Patis    []End     `json:"participant,omitempty"` // 참가자 정보
}

type End struct {
	Type string `json:"type,omitempty"`
	Part string `json:"part,omitempty"`
	Code int    `json:"code,omitempty"`
}

type CallTime struct {
	Start           string `json:"start,omitempty"`
	Answer          string `json:"answer,omitempty"`
	End             string `json:"end,omitempty"`
	ConnectDuration string `json:"connDur,omitempty"`
	TalkDuration    string `json:"talkDur,omitempty"`
}

func main() {
	seg := CallSegment{
		Event:  "test",
		CallID: "asdf",
	}
	b, _ := json.Marshal(seg)
	log.Printf("%v\n", string(b))

	// cfg := Config{}
	// if err := env.Parse(&cfg); err != nil {
	// 	log.Fatalf("%v", err)
	// }

	// if err := mongo.Connect(cfg.MongoServer, cfg.MongoPort); err != nil {
	// 	log.Fatal("Connect MongoDB Fail")
	// }

	// ctx := context.TODO()
	// fmt.Println("connected mongo")
	// fmt.Println("----------------> create")
	// create(ctx)
	// fmt.Println("----------------> create condition")

	// createCondition("test", "1234")
	// fmt.Println("----------------> read all")
	// read()
	// fmt.Println("----------------> read all by id")
	// readCondition("test")
	// fmt.Println("----------------> update")
	// update("test", "kimyeonho")
	// fmt.Println("----------------> delete")
	// delete("test1")
}
