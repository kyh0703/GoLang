package mongo

import (
	"context"
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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	uri := "mongodb://" + host + ":" + port
	clientOpts := options.Client().ApplyURI(uri)
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

}
