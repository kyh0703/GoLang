package transaction

import (
	"framework"
	"log"
	"time"
)

type MakeDid struct {
	framework.Worker
}

func NewMakeDid() *MakeDid {
	return &MakeDid{
		Worker: framework.Worker{
			Id:     "test",
			Type:   MakeDID,
			Expire: time.Second * 30,
		},
	}
}

func (self *MakeDid) Timeout() {
	log.Println("timeout")
}

func (self *MakeDid) DoWork() {
	log.Println("MakeDid")
}
