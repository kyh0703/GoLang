package service

import (
	"context"

	callpb "github.com/kyh0703/grpc/gen/call"
)

type CallService struct {
}

func NewCallService() *CallService {
	return &CallService{}
}

func (s *CallService) Make(ctx context.Context, req *callpb.MakeRequest) (*callpb.MakeResponse, error) {
	return &callpb.MakeResponse{}, nil
}

func (s *CallService) Hold(ctx context.Context, req *callpb.HoldRequest) (*callpb.HoldResponse, error) {
	return &callpb.HoldResponse{}, nil
}

func (s *CallService) Transfer(ctx context.Context, req *callpb.TransferRequest) (*callpb.TransferResponse, error) {
	return &callpb.TransferResponse{}, nil
}

func (s *CallService) MuteTransfer(ctx context.Context, req *callpb.MuteTransferRequest) (*callpb.MuteTransferResponse, error) {
	return &callpb.MuteTransferResponse{}, nil
}

func (s *CallService) Conference(ctx context.Context, req *callpb.ConferenceRequest) (*callpb.ConferenceResponse, error) {
	return &callpb.ConferenceResponse{}, nil
}

func (s *CallService) MuteConference(ctx context.Context, req *callpb.MuteConferenceRequest) (*callpb.MuteConferenceResponse, error) {
	return &callpb.MuteConferenceResponse{}, nil
}

func (s *CallService) Release(ctx context.Context, req *callpb.ReleaseRequest) (*callpb.ReleaseResponse, error) {
	return &callpb.ReleaseResponse{}, nil
}

func (s *CallService) Route(ctx context.Context, req *callpb.RouteRequest) (*callpb.RouteResponse, error) {
	return &callpb.RouteResponse{}, nil
}
