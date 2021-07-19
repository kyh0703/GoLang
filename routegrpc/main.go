package main

import (
	"log"
	"net"

	"github.com/caarlos0/env"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/kyh0703/golang/routegrpc/service"
	routepb "github.com/kyh0703/grpc/gen/route"
	"google.golang.org/grpc"
)

type Config struct {
	Port         string `env:"PORT" envDefault:"9000"`
	IsProduction bool   `env:"PRODUCTION"`
	// SecertKey    string `env:"SECRET_KEY, required"`
}

func main() {
	cfg := Config{}
	env.Parse(&cfg)

	svc := service.NewRouteService()

	svr := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	routepb.RegisterRouteServer(svr, svc)

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatal("failed listening: %s", err)
	}

	if err := svr.Serve(lis); err != nil {
		log.Fatal("server ended: %s", err)
	}
}
