package mongo

import (
	"context"
	"sync"
	"time"

	"github.com/caarlos0/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client      *mongo.Client
	clientError error
	once        sync.Once
)

type Config struct {
	MongoServer string `env:"MONGO_SERVER" envDefault:"127.0.0.1"`
	MongoPort   string `env:"MONGO_PORT" envDefault:"27017"`
	MongoId     string `env:"MONGO_ID" envDefault:"admin"`
	MongoPwd    string `env:"MONGO_PWD" envDefault:"dnflth"`
}

// GetClient is return the connection instance
func GetClient() (*mongo.Client, error) {
	once.Do(func() {
		cfg := &Config{}
		env.Parse(cfg)

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		uri := "mongodb://" + cfg.MongoServer + ":" + cfg.MongoPort
		opts := options.Client().ApplyURI(uri)
		opts.SetMaxPoolSize(100)
		opts.SetMinPoolSize(5)
		opts.SetMaxConnIdleTime(10 * time.Second)

		if cfg.MongoId != "" && cfg.MongoPwd != "" {
			opts.SetAuth(options.Credential{
				Username: cfg.MongoId,
				Password: cfg.MongoPwd,
			})
		}

		client, clientError = mongo.Connect(ctx, opts)
		if clientError != nil {
			return
		}

		clientError = client.Ping(ctx, nil)
		if clientError != nil {
			return
		}
	})

	return client, clientError
}
