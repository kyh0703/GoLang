package main

import (
	"log"
	"net"

	"github.com/caarlos0/env"
	"github.com/kyh0703/golang/routegrpc/mongo"
	"github.com/kyh0703/golang/routegrpc/service"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	callpb "github.com/kyh0703/grpc/gen/call"
)

type Config struct {
	Port string `env:"PORT" envDefault:"9000"`
}

func main() {
	cfg := Config{}
	env.Parse(&cfg)

	if err := mongo.Connect(); err != nil {
		log.Fatalf("init fail mongo conn: %s", err)
	}

	svc := service.NewCallService()
	svr := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)
	callpb.RegisterCallServer(svr, svc)

	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatal("listen port fail")
	}

	if err := svr.Serve(lis); err != nil {
		log.Fatal("grpc server fail")
	}
}
