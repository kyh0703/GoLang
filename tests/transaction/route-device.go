package transaction

import (
	"fmt"
	"framework"
	"log"
	"time"
)

type RouteDevice struct {
	framework.Worker
}

func NewRouteDevice() framework.Work {
	return &RouteDevice{
		Worker: framework.Worker{
			Type:   TypeRouteDevice,
			Expire: time.Second * 30,
		},
	}
}

func (tx *RouteDevice) DoWork() error {
	// 1. Find Device

	// 2. Sip Line New Call 요청

	err := framework.WaitEvent(tx.Context, time.Second*20, func(e framework.Event) error {
		fmt.Println("WaitEvent", e.Id)
		return nil
	}, framework.ExpectEvent{Id: "Leg1", Type: 0})
	if err != nil {
		log.Fatal(err)
		return
	}

	// 3. Event전달
}
