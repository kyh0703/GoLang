package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/caarlos0/env"
	"github.com/kyh0703/golang/routegrpc/service"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	routepb "github.com/kyh0703/grpc/gen/route"
)

type Config struct {
	Port         string `env:"PORT" envDefault:"9000"`
	MongoHost    string `env:"MONGO_SERVER" envDefault: "127.0.0.1"`
	MongoPort    string `env:"MONGO_PORT" envDefault: "27017"`
	IsProduction bool   `env:"PRODUCTION"`
	// SecertKey    string `env:"SECRET_KEY, required"`
}

func mongoConn(cfg Config) (*mongo.Client, error) {
	opts := options.Client().ApplyURI("mongodb://" + cfg.MongoHost + ":" + cfg.MongoPort)
	c, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	if err := c.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	return c, nil
}

func gRPCServer(cfg Config) error {
	svc := service.NewRouteService()
	svr := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)
	routepb.RegisterRouteServer(svr, svc)

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return err
	}

	if err := svr.Serve(lis); err != nil {
		return err
	}

	return nil
}

func main() {
	cfg := Config{}
	env.Parse(&cfg)

	client, err := mongoConn(cfg)
	if err != nil {
		log.Fatalf("init fail mongo conn: %s", err)
	}
	fmt.Println(client)

	if err := gRPCServer(cfg); err != nil {
		log.Fatalf("init fail grpc server: %s", err)
	}
}
