package rpc

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

var kacp = keepalive.ClientParameters{
	Time:                time.Second * 5,
	Timeout:             time.Second,
	PermitWithoutStream: true,
}

func clientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("[c-request] method: %v message: %v\n", method, req)
	if err := invoker(ctx, method, req, reply, cc, opts...); err != nil {
		log.Printf("method: %v, err: %v, req: %v\n", method, err, req)
		return err
	}
	log.Printf("[c-reply] method: %v, message: %v\n", method, reply)
	return nil
}

func Connection(dns string) (*grpc.ClientConn, error) {
	resolver.SetDefaultScheme("dns")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithKeepaliveParams(kacp))
	opts = append(opts, grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			clientInterceptor,
		),
	))
	hostName, _ := os.Hostname()
	opts = append(opts, grpc.WithUserAgent(hostName))
	return grpc.Dial(dns, opts...)
}
