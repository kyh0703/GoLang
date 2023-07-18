package service

import (
	"fmt"
	"os"
	"strings"

	platform_server "gitlab.com/ipron-cloud/iproncloud-platform-server"
	service_data "gitlab.com/ipron-cloud/iproncloud-platform-server/archapi"
)

const (
	InsertHook    = "insert"
	UpdateHook    = "update"
	DeleteHook    = "delete"
	RefreshHook   = "refresh"
	RefreshedHook = "refreshed"
)

const DataStreamTopicName = "bridgetec.ipron.dbstream"

type HookHandler interface {
	onInsert(id string, data interface{}) error
	onUpdate(id string, data interface{}) error
	onDelete(id string)
}

var Hook = make(HookManager)

type HookManager map[string]HookHandler

func (m HookManager) Collections() string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return strings.Join(keys, ",")
}

func (m HookManager) GetHandler(collection string) (HookHandler, bool) {
	v, ok := m[collection]
	return v, ok
}

func StartHook() {
	p := platform_server.New()
	var (
		service = p.ArchApi()
		options []service_data.Options
	)

	// TODO remove TestCode
	service.Connect("100.100.103.163:6001")

	hostname, _ := os.Hostname()
	options = append(options, service_data.WithLocalCacheOption("100.100.103.163:31609", hostname, DataStreamTopicName, Hook.Collections()))
	options = append(options, service_data.WithDataHookOption(onStreamHook))
	service.LocalCacheStart(options...)
}

func onStreamHook(action, collection, tenant, id string, data interface{}) {
	handler, _ := Hook.GetHandler(collection)
	switch action {
	case InsertHook:
		handler.onInsert(id, data)
	case UpdateHook:
		handler.onUpdate(id, data)
	case DeleteHook:
		handler.onDelete(id)
	case RefreshHook:
		fmt.Println("refresh - ", collection)
	case RefreshedHook:
		fmt.Println("refreshed - ", collection)
	}
}
