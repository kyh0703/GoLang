package config

import (
	"log"
	"time"

	"github.com/caarlos0/env"
)

var Env EnvConfig

type EnvConfig struct {
	PodNamespace         string        `env:"MY_POD_NAMESPACE" envDefault:"192.168.115.148"`
	PodName              string        `env:"MY_POD_NAME" envDefault:"192.168.115.148"`
	PodIP                string        `env:"MY_POD_IP" envDefault:"192.168.115.148"`
	Port                 string        `env:"GRPC_PORT" envDefault:"8090"`
	ReadTimeout          time.Duration `env:"READ_TIMEOUT" envDefault: 60s`
	JaegerControllerDNS  string        `env:"BCLOUD_TRACES_URI" envDefault:"http://jaeger-collector.istio-system:14268/api/traces"`
	KafkaBroker          []string      `env:"BCLOUD_KAFKA_BROKER_URI" envSeparator:"," envDefault:"100.100.103.163:31090"`
	QueueSvcDNS          string        `env:"BCLOUD_QFLOW_SVC_URI" envDefault:"100.100.103.163:50052"`
	CallFlowSvcDNS       string        `env:"BCLOUD_CALLFLOW_SVC_URI" envDefault:"100.100.103.163:50052"`
	CompatibleFlowSvcDNS string        `env:"BCLOUD_COMPATIBLE_FLOW_SVC_URI" envDefault:"100.100.103.163:50052"`
	ArchDataSvcDNS       string        `env:"BCLOUD_ARCHDATA_SVC_URI" envDefault:"100.100.103.163:6001"`
	PresenceSvcDNS       string        `env:"BCLOUD_PRESENCE_SVC_URI" envDefault:"100.100.103.163:6004"`
	RealtimeSvcDNS       string        `env:"BCLOUD_REALTIMEDATA_SVC_URI" envDefault:"100.100.103.163:6003"`
	EventNotifySvcDNS    string        `env:"BCLOUD_EVTNOTIFY_SVC_URI" envDefault:"100.100.103.163:6002"`
	RecordGatewaySvcDNS  string        `env:"BCLOUD_RECORD_GATEWAY_SVC_URI" envDefault:"recgw-svc:80"`
	TrunkSvcDNS          string        `env:"BCLOUD_TRUNK_SVC_URI" envDefault:"trunksip:80"`
}

func init() {
	env.Parse(&Env)
	log.Printf("--------- Environment ---------\n")
	log.Printf("MY_POD_NAMESPACE               : %v\n", Env.PodNamespace)
	log.Printf("MY_POD_NAME                    : %v\n", Env.PodName)
	log.Printf("MY_POD_IP                      : %v\n", Env.PodIP)
	log.Printf("BCLOUD_TRACES_URI              : %v\n", Env.JaegerControllerDNS)
	log.Printf("BCLOUD_KAFKA_BROKER_DNS        : %v\n", Env.KafkaBroker)
	log.Printf("BCLOUD_QFLOW_SVC_DNS           : %v\n", Env.QueueSvcDNS)
	log.Printf("BCLOUD_CALLFLOW_SVC_DNS        : %v\n", Env.CallFlowSvcDNS)
	log.Printf("BCLOUD_COMPATIBLE_FLOW_SVC_DNS : %v\n", Env.CompatibleFlowSvcDNS)
	log.Printf("BCLOUD_ARCHDATA_SVC_DNS        : %v\n", Env.ArchDataSvcDNS)
	log.Printf("BCLOUD_PRESENCE_SVC_DNS        : %v\n", Env.PresenceSvcDNS)
	log.Printf("BCLOUD_REALTIMEDATA_SVC_DNS    : %v\n", Env.RealtimeSvcDNS)
	log.Printf("BCLOUD_EVTNOTIFY_SVC_DNS       : %v\n", Env.EventNotifySvcDNS)
	log.Printf("BCLOUD_RECORD_GATEWAY_SVC_DNS  : %v\n", Env.RecordGatewaySvcDNS)
	log.Printf("BCLOUD_TRUNK_SVC_DNS           : %v\n", Env.TrunkSvcDNS)
	log.Printf("-------------------------------\n")
}
