package rpc

import (
	"context"
	"errors"

	"gitlab.com/ipron-cloud/call-sdk/config"

	eventpb "gitlab.com/ipron-cloud/grpc-idl/gen/go/event"
	flowpb "gitlab.com/ipron-cloud/grpc-idl/gen/go/flow"
	mediapb "gitlab.com/ipron-cloud/grpc-idl/gen/go/media"
	queuepb "gitlab.com/ipron-cloud/grpc-idl/gen/go/queueflow"
	recordpb "gitlab.com/ipron-cloud/grpc-idl/gen/go/recgw"
	sippb "gitlab.com/ipron-cloud/grpc-idl/gen/go/sip"
)

var (
	errEmptyURI      = errors.New("empty URI")
	errEmptyInstance = errors.New("gRPC instance nil")
)

// ANCHOR Record GW
func StartRecord(ctx context.Context, req *recordpb.RecStartReq) (*recordpb.RecStartRes, error) {
	if recordgw == nil {
		return nil, errEmptyInstance
	}
	return recordgw.RecordStart(ctx, req)
}

func StopRecord(ctx context.Context, req *recordpb.RecStopReq) (*recordpb.RecStopResp, error) {
	if recordgw == nil {
		return nil, errEmptyInstance
	}
	return recordgw.RecordStop(ctx, req)
}

// ANCHOR Trunk API
func TrunkRouteCall(ctx context.Context, req *sippb.RoutecallReq) (*sippb.RoutecallResp, error) {
	if trunk == nil {
		return nil, errEmptyInstance
	}
	return trunk.RouteCall(ctx, req)
}

// ANCHOR SIP SERVICE
func GetConnCountPhone(ctx context.Context, uri string, req *sippb.PhoneConnCountRequest) (*sippb.PhoneConnCountReply, error) {
	if len(uri) == 0 {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return sippb.NewIpronSipClient(gc).PhoneConnCount(ctx, req)
}

func SipNewCall(ctx context.Context, uri string, req *sippb.NewcallReq) (*sippb.NewcallResp, error) {
	if len(uri) == 0 {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return sippb.NewIpronSipClient(gc).NewCall(ctx, req)
}

func SipRelease(ctx context.Context, uri string, req *sippb.ReleaseReq) (*sippb.ReleaseResp, error) {
	if len(uri) == 0 {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return sippb.NewIpronSipClient(gc).Release(ctx, req)
}

func SipHeld(ctx context.Context, uri string, req *sippb.HeldReq) (*sippb.HeldResp, error) {
	if len(uri) == 0 {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return sippb.NewIpronSipClient(gc).Held(ctx, req)
}

func SipRetrive(ctx context.Context, uri string, req *sippb.RetrieveReq) (*sippb.RetrieveResp, error) {
	if len(uri) == 0 {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return sippb.NewIpronSipClient(gc).Retrieve(ctx, req)
}

func SipNotifyAnswer(ctx context.Context, uri string, req *sippb.NotifyTalkReq) (*sippb.NotifyTalkResp, error) {
	if len(uri) == 0 {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return sippb.NewIpronSipClient(gc).NotifyTalk(ctx, req)
}

func SipNotifyHold(ctx context.Context, uri string, req *sippb.NotifyHoldReq) (*sippb.NotifyHoldResp, error) {
	if len(uri) == 0 {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return sippb.NewIpronSipClient(gc).NotifyHold(ctx, req)
}

func SipTransfer(ctx context.Context, uri string, req *sippb.TransferReq) (*sippb.TransferResp, error) {
	if len(uri) == 0 {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return sippb.NewIpronSipClient(gc).Transfer(ctx, req)
}

func SipNotifyTransfer(ctx context.Context, uri string, req *sippb.NotifyTransferReq) (*sippb.NotifyTransferResp, error) {
	if len(uri) == 0 {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return sippb.NewIpronSipClient(gc).NotifyTransfer(ctx, req)
}

func SipConference(ctx context.Context, uri string, req *sippb.ConfJoinReq) (*sippb.ConfJoinResp, error) {
	if len(uri) == 0 {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return sippb.NewIpronSipClient(gc).ConfJoin(ctx, req)
}

// ANCHOR MEDIA SERVICE
func PlayMedia(ctx context.Context, uri string, req *mediapb.PlayReq) (*mediapb.PlayResp, error) {
	if len(uri) == 0 {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return mediapb.NewMediaClient(gc).Play(ctx, req)
}

func StopMedia(ctx context.Context, uri string, req *mediapb.StopReq) (*mediapb.StopResp, error) {
	if len(uri) == 0 {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return mediapb.NewMediaClient(gc).Stop(ctx, req)
}

func Record(ctx context.Context, uri string, req *mediapb.RecordReq) (*mediapb.RecordResp, error) {
	if len(uri) == 0 {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return mediapb.NewMediaClient(gc).Record(ctx, req)
}

// ANCHOR Compatible Flow Service
func CompatibleFlowStart(ctx context.Context, req *flowpb.FlowStartCallReq) (*flowpb.FlowStartCallRes, error) {
	if compatibleflow == nil {
		return nil, errEmptyInstance
	}
	return compatibleflow.FlowStartCall(ctx, req)
}

// ANCHOR FLOW SERVICE
func FlowStart(ctx context.Context, req *flowpb.FlowStartCallReq) (*flowpb.FlowStartCallRes, error) {
	if flow == nil {
		return nil, errEmptyInstance
	}
	return flow.FlowStartCall(ctx, req)
}

// ANCHOR QUEUE SERVICE
func QueueFlowStart(ctx context.Context, req *queuepb.QueueFlowStartReq) (*queuepb.QueueFlowStartRes, error) {
	if queue == nil {
		return nil, errEmptyInstance
	}
	return queue.QueueFlowStart(ctx, req)
}

// ANCHOR EVENT SERVICE
func EventAlerted(ctx context.Context, uri string, req *eventpb.EventAlertedRequest) (*eventpb.EventAlertedReply, error) {
	if len(uri) == 0 || uri == config.Env.PodIP {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return eventpb.NewEventClient(gc).EventAlerted(ctx, req)
}

func EventConnected(ctx context.Context, uri string, req *eventpb.EventConnectedRequest) (*eventpb.EventConnectedReply, error) {
	if len(uri) == 0 || uri == config.Env.PodIP {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return eventpb.NewEventClient(gc).EventConnected(ctx, req)
}

func EventHeld(ctx context.Context, uri string, req *eventpb.EventHeldRequest) (*eventpb.EventHeldReply, error) {
	if len(uri) == 0 || uri == config.Env.PodIP {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return eventpb.NewEventClient(gc).EventHeld(ctx, req)
}

func EventRetrived(ctx context.Context, uri string, req *eventpb.EventRetrivedRequest) (*eventpb.EventRetrivedReply, error) {
	if len(uri) == 0 || uri == config.Env.PodIP {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return eventpb.NewEventClient(gc).EventRetrived(ctx, req)
}

func EventTransfered(ctx context.Context, uri string, req *eventpb.EventTransferedRequest) (*eventpb.EventTransferedReply, error) {
	if len(uri) == 0 || uri == config.Env.PodIP {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return eventpb.NewEventClient(gc).EventTransfered(ctx, req)
}

func EventConference(ctx context.Context, uri string, req *eventpb.EventConferenceRequest) (*eventpb.EventConferenceReply, error) {
	if len(uri) == 0 || uri == config.Env.PodIP {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return eventpb.NewEventClient(gc).EventConference(ctx, req)
}

func EventReleased(ctx context.Context, uri string, req *eventpb.EventReleasedRequest) (*eventpb.EventReleasedReply, error) {
	if len(uri) == 0 || uri == config.Env.PodIP {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return eventpb.NewEventClient(gc).EventReleased(ctx, req)
}

func EventRouteTimeout(ctx context.Context, uri string, req *eventpb.EventRouteTimeoutRequest) (*eventpb.EventRouteTimeoutReply, error) {
	if len(uri) == 0 || uri == config.Env.PodIP {
		return nil, errEmptyURI
	}
	gc, err := Connection(uri)
	if err != nil {
		return nil, err
	}
	defer gc.Close()
	return eventpb.NewEventClient(gc).EventRouteTimeout(ctx, req)
}
