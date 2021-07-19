package service

import (
	"context"

	routepb "github.com/kyh0703/grpc/gen/route"
)

type RouteService struct {
}

func NewRouteService() *RouteService {
	return &RouteService{}
}

func (s *RouteService) Device(ctx context.Context, req *routepb.DeviceRequest) (*routepb.DeviceResponse, error) {

	return &routepb.DeviceResponse{}, nil
}

func (s *RouteService) Agent(ctx context.Context, req *routepb.AgentRequest) (*routepb.AgentResponse, error) {

	return &routepb.AgentResponse{}, nil
}

func (s *RouteService) Trunk(ctx context.Context, req *routepb.TrunkRequest) (*routepb.TrunkResponse, error) {

	return &routepb.TrunkResponse{}, nil
}

func (s *RouteService) Ring(ctx context.Context, req *routepb.RingRequest) (*routepb.RingResponse, error) {

	return &routepb.RingResponse{}, nil
}

func (s *RouteService) Answer(ctx context.Context, req *routepb.AnswerRequest) (*routepb.AnswerResponse, error) {

	return &routepb.AnswerResponse{}, nil
}
