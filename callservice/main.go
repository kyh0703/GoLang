package main

import (
	"log"
	"net"

	"github.com/caarlos0/env"
	"github.com/kyh0703/golang/callservice/service"
	"google.golang.org/grpc"

	pb "git.bridgetec.com/IPRON/idl.git/gen/go"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

type Config struct {
	Port string `env:"PORT" envDefault:"9000"`
}

func main() {
	cfg := Config{}
	env.Parse(&cfg)

	svc := service.NewCallService()
	svr := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)
	pb.RegisterCallServer(svr, svc)

	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatal("listen port fail")
	}

	if err := svr.Serve(lis); err != nil {
		log.Fatal("grpc server fail")
	}
}
