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
			Id:     "test",
			Type:   Make3PCC,
			Expire: time.Second * 3,
		},
	}
}

func (tx *Make3pcc) DoWork() {
	log.Printf("[%v] Make3pcc", tx.Id)
	select {
	case <-tx.GetContext().Done():
		fmt.Printf("context done\n")
		return
	}
}
