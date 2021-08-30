package transaction

import (
	"fmt"
	"framework"
	"log"
	"time"
)

type Make3pcc struct {
	framework.Worker
}

func NewMake3pcc() framework.Work {
	return &Make3pcc{
		Worker: framework.Worker{
			Type:   TypeMake3pcc,
			Expire: time.Second * 30,
		},
	}
}

func Test(e framework.Event) error {
	return nil
}

func (tx *Make3pcc) DoWork() {
	err := framework.WaitEvent(tx.Context, time.Second*20, func(e framework.Event) error {
		fmt.Println("WaitEvent", e.Id)
		return nil
	}, framework.ExpectEvent{Id: "Leg1", Type: 0})
	if err != nil {
		log.Fatal(err)
		return
	}

	// work, err := framework.MakeWorker(TypeRouteDevice)
	// if err != nil {
	// 	return
	// }
	// work.doWork()

}
