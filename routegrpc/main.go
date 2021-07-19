package main

import (
	"log"

	"github.com/caarlos0/env"

	"github.com/golang/protobuf/protoc-gen-go/grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	routepb "github.com/kyh0703/grpc/gen/route"
	"github.com/kyh0703/routegrpc/service"
)

type Config struct {
	Port         string `env:"PORT" envDefault:"9000"`
	IsProduction bool   `env:"PRODUCTION"`
	// SecertKey    string `env:"SECRET_KEY, required"`
}

// func NewServer(cfg Config) error {
// 	lis, err := net.Listen("tcp", ":"+cfg.Port)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 		return err
// 	}

// 	log.Printf("start gRPC server on %s port", cfg.Port)
// 	return serve.Serve(lis)
// }

func main() {
	cfg := Config{}
	env.Parse(&cfg)

	if err := NewServer(cfg); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	svc := service.NewRouteService()

	svr := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	routepb.RegisterRouteServer(svr, svc)
}
