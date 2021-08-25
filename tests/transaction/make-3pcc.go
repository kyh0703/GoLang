package transaction

import (
	"framework"
	"log"
	"time"
)

type Make3pcc struct {
	framework.Worker
}

func NewMake3pcc() *Make3pcc {
	return &Make3pcc{
		Worker: framework.Worker{
			Id:     "test",
			Type:   Make3PCC,
			Expire: time.Second * 30,
		},
	}
}

func (self *Make3pcc) Timeout() {
	log.Println("timeout")
}

func (self *Make3pcc) DoWork() {
	log.Println("Make3pcc")
}
