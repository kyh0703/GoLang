package route

import (
	"context"

	routepb "github.com/kyh0703/grpc/gen/route"
)

type router struct {
	routepb.RouteServer
}

func (s *router) Device(ctx context.Context, req *routepb.DeviceRequest) (*routepb.DeviceResponse, error) {

	return &routepb.DeviceResponse{}, nil
}

func (s *router) Agent(ctx context.Context, req *routepb.AgentRequest) (*routepb.AgentResponse, error) {

	return &routepb.AgentResponse{}, nil
}

func (s *router) Trunk(ctx context.Context, req *routepb.TrunkRequest) (*routepb.TrunkResponse, error) {

	return &routepb.TrunkResponse{}, nil
}

func (s *router) Ring(ctx context.Context, req *routepb.RingRequest) (*routepb.RingResponse, error) {

	return &routepb.RingResponse{}, nil
}

func (s *router) Answer(ctx context.Context, req *routepb.AnswerRequest) (*routepb.AnswerResponse, error) {

	return &routepb.AnswerResponse{}, nil
}
