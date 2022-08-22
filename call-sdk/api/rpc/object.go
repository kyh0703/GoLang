package rpc

import (
	"sync"

	"gitlab.com/ipron-cloud/call-app/pkg/config"
	"google.golang.org/grpc"

	archpb "gitlab.com/ipron-cloud/grpc-idl/gen/go/archdata"
	flowpb "gitlab.com/ipron-cloud/grpc-idl/gen/go/flow"
	queuepb "gitlab.com/ipron-cloud/grpc-idl/gen/go/queueflow"
	recordpb "gitlab.com/ipron-cloud/grpc-idl/gen/go/recgw"
	sippb "gitlab.com/ipron-cloud/grpc-idl/gen/go/sip"
)

var (
	once           sync.Once
	flow           flowpb.FlowClient
	queue          queuepb.QueueFlowClient
	arch           archpb.ArchDataClient
	trunk          sippb.IpronSipClient
	compatibleflow flowpb.FlowClient
	recordgw       recordpb.RecDeliverServiceClient
)

func init() {
	once.Do(func() {
		var gc *grpc.ClientConn
		gc, _ = Connection(config.Env.CallFlowSvcDNS)
		flow = flowpb.NewFlowClient(gc)
		gc, _ = Connection(config.Env.QueueSvcDNS)
		queue = queuepb.NewQueueFlowClient(gc)
		gc, _ = Connection(config.Env.ArchDataSvcDNS)
		arch = archpb.NewArchDataClient(gc)
		gc, _ = Connection(config.Env.TrunkSvcDNS)
		trunk = sippb.NewIpronSipClient(gc)
		gc, _ = Connection(config.Env.CompatibleFlowSvcDNS)
		compatibleflow = flowpb.NewFlowClient(gc)
		gc, _ = Connection(config.Env.RecordGatewaySvcDNS)
		recordgw = recordpb.NewRecDeliverServiceClient(gc)
	})
}
