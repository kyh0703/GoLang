package service

import (
	"sync"

	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

var (
	once sync.Once
	cli  *grpc.ClientConn
)

func GetClient(servicehost string) *grpc.ClientConn {
	once.Do(func() {
		conn, _ := grpc.Dial(
			servicehost,
			grpc.WithUnaryInterceptor(
				grpc_middleware.ChainUnaryClient(),
			),
		)

		cli = conn
	})

	return cli
}
