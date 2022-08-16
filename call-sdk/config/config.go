package config

import (
	"time"

	"github.com/caarlos0/env"
)

var (
	Env          EnvConfig
	LocalAddress string
)

type EnvConfig struct {
	PodNamespace         string        `env:"MY_POD_NAMESPACE" envDefault:"192.168.115.148"`
	PodName              string        `env:"MY_POD_NAME" envDefault:"192.168.115.148"`
	PodIP                string        `env:"MY_POD_IP" envDefault:"192.168.115.148"`
	Port                 string        `env:"PORT" envDefault:"8090"`
	PlayMentTimeout      time.Duration `env:"PLAY_MENT_TIMEOUT" envDefault:"60s"`
	RingWaitTimeout      time.Duration `env:"RING_WAIN_TIMEOUT" envDefault:"180s"`
	TransactionTimeout   time.Duration `env:"TRANSACTION_TIMEOUT" envDefault:"180s"`
	JaegerControllerDNS  string        `env:"BCLOUD_TRACES_URI" envDefault:"http://jaeger-collector.istio-system:14268/api/traces"`
	KafkaBroker          []string      `env:"BCLOUD_KAFKA_BROKER_URI" envSeparator:"," envDefault:"100.100.103.163:31090"`
	QueueSvcDNS          string        `env:"BCLOUD_QFLOW_SVC_URI" envDefault:"100.100.103.163:50052"`
	CallFlowSvcDNS       string        `env:"BCLOUD_CALLFLOW_SVC_URI" envDefault:"100.100.103.163:50052"`
	CompatibleFlowSvcDNS string        `env:"BCLOUD_COMPATIBLE_FLOW_SVC_URI" envDefault:"100.100.103.163:50052"`
	ArchDataSvcDNS       string        `env:"BCLOUD_ARCHDATA_SVC_URI" envDefault:"archdatasvc:80"`
	PresenceSvcDNS       string        `env:"BCLOUD_PRESENCE_SVC_URI" envDefault:"100.100.103.163:30509"`
	RealtimeSvcDNS       string        `env:"BCLOUD_REALTIMEDATA_SVC_URI" envDefault:"100.100.103.163:30511"`
	EventNotifySvcDNS    string        `env:"BCLOUD_EVTNOTIFY_SVC_URI" envDefault:"100.100.103.163:6002"`
	RecordGatewaySvcDNS  string        `env:"BCLOUD_RECORD_GATEWAY_SVC_URI" envDefault:"recgw-svc:80"`
	TrunkSvcDNS          string        `env:"BCLOUD_TRUNK_SVC_URI" envDefault:"trunksip:80"`
}

func init() {
	env.Parse(&Env)
	LocalAddress = Env.PodIP + ":" + Env.Port
	log.Info("--------- Environment ---------")
	log.Info("MY_POD_NAMESPACE               : %v", Env.PodNamespace)
	log.Info("MY_POD_NAME                    : %v", Env.PodName)
	log.Info("MY_POD_IP                      : %v", Env.PodIP)
	log.Info("PLAY_MENT_TIMEOUT              : %v", Env.PlayMentTimeout)
	log.Info("RING_WAIN_TIMEOUT              : %v", Env.RingWaitTimeout)
	log.Info("TRANSACTION_TIMEOUT            : %v", Env.TransactionTimeout)
	log.Info("BCLOUD_TRACES_URI              : %v", Env.JaegerControllerDNS)
	log.Info("BCLOUD_KAFKA_BROKER_DNS        : %v", Env.KafkaBroker)
	log.Info("BCLOUD_QFLOW_SVC_DNS           : %v", Env.QueueSvcDNS)
	log.Info("BCLOUD_CALLFLOW_SVC_DNS        : %v", Env.CallFlowSvcDNS)
	log.Info("BCLOUD_COMPATIBLE_FLOW_SVC_DNS : %v", Env.CompatibleFlowSvcDNS)
	log.Info("BCLOUD_ARCHDATA_SVC_DNS        : %v", Env.ArchDataSvcDNS)
	log.Info("BCLOUD_PRESENCE_SVC_DNS        : %v", Env.PresenceSvcDNS)
	log.Info("BCLOUD_REALTIMEDATA_SVC_DNS    : %v", Env.RealtimeSvcDNS)
	log.Info("BCLOUD_EVTNOTIFY_SVC_DNS       : %v", Env.EventNotifySvcDNS)
	log.Info("BCLOUD_RECORD_GATEWAY_SVC_DNS  : %v", Env.RecordGatewaySvcDNS)
	log.Info("BCLOUD_TRUNK_SVC_DNS           : %v", Env.TrunkSvcDNS)
	log.Info("-------------------------------")
}
