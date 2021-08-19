package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	pb "git.bridgetec.com/IPRON/idl.git/gen/go"
)

const address = "localhost:9000"

func TestMakeBySipTrunk(t *testing.T) {
	assert := assert.New(t)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	assert.Nilf(err, "did not connect: %v", err)
	defer conn.Close()

	client := pb.NewCallClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t.Run("Make", func(t *testing.T) {
		res, err := client.Make(ctx, &pb.MakeRequest{})
		assert.Nilf(err, "did not connect: %v", err)
		t.Logf("Response: %v", res)
	})
}

func TestMakeBySipLine(t *testing.T) {
	assert := assert.New(t)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	assert.Nilf(err, "did not connect :%v", err)
	defer conn.Close()

	client := pb.NewCallClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t.Run("Make", func(t *testing.T) {
		res, err := client.Make(ctx, &pb.MakeRequest{})
		assert.Nil(err)
		t.Logf("Response: %v", res)
	})
}
