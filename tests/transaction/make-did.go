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
			Type:   TypeMakeDid,
			Expire: time.Second * 30,
		},
	}
}

func (self *MakeDid) DoWork() error {
	log.Printf("[%v] MakeDid", self.Id)
}
