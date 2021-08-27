package framework

import (
	"context"
	"errors"
	"time"
)

type WorkerType uint8

type Work interface {
	GetId() string
	GetType() WorkerType
	GetExpire() time.Duration
	GetContext() context.Context
	GetCancelFunc() context.CancelFunc
	SetId(id string)
	SetType(kind WorkerType)
	SetExpire(expire time.Duration)
	SetContext(ctx context.Context)
	SetCancelFunc(cf context.CancelFunc)
	// below is Virtual Functions
	DoWork()
}

type Worker struct {
	Id     string
	Type   WorkerType
	Expire time.Duration
	context.Context
	context.CancelFunc
}

func (w *Worker) GetId() string {
	return w.Id
}

func (w *Worker) GetType() WorkerType {
	return w.Type
}

func (w *Worker) GetExpire() time.Duration {
	return w.Expire
}

func (w *Worker) GetContext() context.Context {
	return w.Context
}

func (w *Worker) GetCancelFunc() context.CancelFunc {
	return w.CancelFunc
}

func (w *Worker) SetId(id string) {
	w.Id = id
}

func (w *Worker) SetType(t WorkerType) {
	w.Type = t
}

func (w *Worker) SetExpire(expire time.Duration) {
	w.Expire = expire
}

func (w *Worker) SetContext(ctx context.Context) {
	w.Context = ctx
}

func (w *Worker) SetCancelFunc(cf context.CancelFunc) {
	w.CancelFunc = cf
}

type Workers map[string]Work

var (
	errNotFound     = errors.New("Not Found")
	errCantUpdate   = errors.New("Cant update non-exiting word")
	errWorkerExists = errors.New("That worker already exists")
)

func (w Workers) search(id string) (Work, error) {
	Work, exist := w[id]
	if exist {
		return Work, nil
	}
	return nil, errNotFound
}

func (w Workers) add(id string, worker Work) error {
	_, err := w.search(id)
	switch err {
	case errNotFound:
		w[id] = worker
	case nil:
		return errWorkerExists
	}
	return nil
}

func (w Workers) delete(id string) error {
	_, err := w.search(id)
	if err != nil {
		return errNotFound
	}
	delete(w, id)
	return nil
}

func (w Workers) clear() {
	for k := range w {
		delete(w, k)
	}
}
