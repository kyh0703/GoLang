package main

import (
	"framework"
	"time"
	"transaction"
)

func main() {
	manager := framework.NewManager(transaction.MakeTransaction)
	manager.BeginWorker("12341234", transaction.TypeMake3pcc)

	timer := time.NewTimer(time.Second * 2)
	go func() {
		<-timer.C
		evt := &framework.Event{Id: "Leg1", Type: 0, Data: "123123"}
		manager.Emit("12341234", evt)
	}()

	manager.PrintWorker()
	manager.WaitGroup.Wait()
}
