package transaction

import (
	"framework"
)

const (
	Make3PCC = iota
	MakeDID
)

func MakeTransaction(workerType framework.WorkerType) *framework.Work {
	var worker framework.Work = nil
	switch workerType {
	case Make3PCC:
		worker = NewMake3pcc()
	case MakeDID:
		worker = NewMakeDid()
	}
	return &worker
}
