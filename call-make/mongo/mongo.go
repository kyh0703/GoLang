package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/caarlos0/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Config struct {
	MongoServer string `env:"MONGO_SERVER" envDefault:"127.0.0.1"`
	MongoPort   string `env:"MONGO_PORT" envDefault:"27017"`
	MongoId     string `env:"MONGO_ID" envDefault:"admin"`
	MongoPwd    string `env:"MONGO_PWD" envDefault:"dnflth"`
}

// GetClient is return the connection in ConnectionPool
func GetClient() *mongo.Client {
	return client
}

// Connect MongoDB
func Connect() error {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return errors.New("parsing Env Fail")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := "mongodb://" + cfg.MongoServer + ":" + cfg.MongoPort
	clientOpts := options.Client().ApplyURI(uri)
	clientOpts.SetMaxPoolSize(100)
	clientOpts.SetMinPoolSize(5)
	clientOpts.SetMaxConnIdleTime(10 * time.Second)

	conn, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return err
	}

	err = conn.Ping(ctx, nil)
	if err != nil {
		return err
	}

	client = conn
	return nil
}

// Connect mongoDB with authentication
func ConnectAuth() error {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return errors.New("parsing Env Fail")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := "mongodb://" + cfg.MongoServer + ":" + cfg.MongoPort
	clientOpts := options.Client().ApplyURI(uri)
	clientOpts.SetAuth(options.Credential{
		Username: cfg.MongoId,
		Password: cfg.MongoPwd,
	})
	clientOpts.SetMaxPoolSize(100)
	clientOpts.SetMinPoolSize(10)
	clientOpts.SetMaxConnIdleTime(10 * time.Second)

	conn, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return err
	}

	err = conn.Ping(ctx, nil)
	if err != nil {
		return err
	}

	client = conn
	return nil
}
