package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// GetClient is return the connection in ConnectionPool
func GetClient() *mongo.Client {
	return client
}

// Connect MongoDB
func Connect(host, port string) error {
	fmt.Println("hihihihi")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	uri := "mongodb://" + host + ":" + port
	clientOpts := options.Client().ApplyURI(uri)

	if err := connect(ctx, clientOpts); err != nil {
		return err
	}

	return nil
}

// Connect mongoDB with authentication
func ConnectAuth(host, port, id, pwd string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	uri := "mongodb://" + host + ":" + port
	clientOpts := options.Client().ApplyURI(uri)
	clientOpts.SetAuth(options.Credential{
		Username: id,
		Password: pwd,
	})

	if err := connect(ctx, clientOpts); err != nil {
		return err
	}

	return nil
}

// connect mongoDB
func connect(ctx context.Context, opts *options.ClientOptions) error {
	opts.SetMaxPoolSize(100)
	opts.SetMinPoolSize(10)
	opts.SetMaxConnIdleTime(10 * time.Second)

	conn, err := mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}

	err = conn.Ping(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
