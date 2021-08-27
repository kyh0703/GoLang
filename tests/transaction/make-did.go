package transaction

import (
	"framework"
	"log"
	"time"
)

type MakeDid struct {
	framework.Worker
}

func NewMakeDid() framework.Work {
	return &MakeDid{
		Worker: framework.Worker{
			Id:     "test",
			Type:   MakeDID,
			Expire: time.Second * 30,
		},
	}
}

func (self *MakeDid) DoWork() {
	log.Printf("[%v] MakeDid", self.Id)
}
