package main

import (
	"fmt"
	"framework"
	"transaction"
)

func main() {
	factory := &transaction.Factory{}
	manager := framework.NewManager(factory)
	manager.PrintWorker()

	var w framework.Work = transaction.NewMake3pcc()
	// convert(w)
	fmt.Println(w.GetId())
	w.DoWork()
}
