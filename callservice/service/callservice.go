package service

import (
	"context"

	callpb "git.bridgetec.com/IPRON/idl.git/gen/go"
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

func (s *CallService) Retrive(ctx context.Context, req *callpb.RetriveRequest) (*callpb.RetriveResponse, error) {
	return &callpb.RetriveResponse{}, nil
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

func (s *CallService) EventAlerted(ctx context.Context, req *callpb.EventAlertedRequest) (*callpb.EventAlertedResponse, error) {
	return &callpb.EventAlertedResponse{}, nil
}

func (s *CallService) EventHeld(ctx context.Context, req *callpb.EventHeldRequest) (*callpb.EventHeldResponse, error) {
	return &callpb.EventHeldResponse{}, nil
}

func (s *CallService) EventRetrived(ctx context.Context, req *callpb.EventRetrivedRequest) (*callpb.EventRetrivedResponse, error) {
	return &callpb.EventRetrivedResponse{}, nil
}

func (s *CallService) EventTransfered(ctx context.Context, req *callpb.EventTransferedRequest) (*callpb.EventTransferedResponse, error) {
	return &callpb.EventTransferedResponse{}, nil
}

func (s *CallService) EventConferenced(ctx context.Context, req *callpb.EventConferencedRequest) (*callpb.EventConferencedResponse, error) {
	return &callpb.EventConferencedResponse{}, nil
}

func (s *CallService) EventConnected(ctx context.Context, req *callpb.EventConnectedRequest) (*callpb.EventConnectedResponse, error) {
	return &callpb.EventConnectedResponse{}, nil
}

func (s *CallService) EventReleased(ctx context.Context, req *callpb.EventReleasedRequest) (*callpb.EventReleasedResponse, error) {
	return &callpb.EventReleasedResponse{}, nil
}
