package transaction

import (
	"errors"
	"framework"
)

var errTypeNotFound = errors.New("Invalid worker type")

const (
	TypeMake3pcc = iota
	TypeMakeDid
	TypeRouteDevice
)

func MakeTransaction(workerType framework.WorkerType) (framework.Work, error) {
	switch workerType {
	case TypeMake3pcc:
		return NewMake3pcc(), nil
	case TypeMakeDid:
		return NewMakeDid(), nil
	case TypeRouteDevice:
		return NewRouteDevice(), nil
	default:
		return nil, errTypeNotFound
	}
}
