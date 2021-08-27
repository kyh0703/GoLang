package transaction

import (
	"errors"
	"framework"
)

var errTypeNotFound = errors.New("Invalid worker type")

const (
	Make3PCC = iota
	MakeDID
)

func MakeTransaction(workerType framework.WorkerType) (framework.Work, error) {
	switch workerType {
	case Make3PCC:
		return NewMake3pcc(), nil
	case MakeDID:
		return NewMakeDid(), nil
	default:
		return nil, errTypeNotFound
	}
}
