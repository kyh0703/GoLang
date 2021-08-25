package framework

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type WorkerType uint8

type Work interface {
	GetId() string
	GetType() WorkerType
	GetExpire() time.Duration
	SetId(id string)
	SetType(kind WorkerType)
	SetExpire(expire time.Duration)
	// below is Virtual Functions
	DoWork()
	Timeout()
}

type Worker struct {
	Id         string
	Type       WorkerType
	Expire     time.Duration
	Ctx        context.Context
	CancelFunc context.CancelFunc
	Events     chan Event
	keepEvent  []Event
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

func (w *Worker) SetId(id string) {
	w.Id = id
}

func (w *Worker) SetType(t WorkerType) {
	w.Type = t
}

func (w *Worker) SetExpire(expire time.Duration) {
	w.Expire = expire
}

func (w *Worker) Emmit(event Event) {
	w.Events <- event
}

func (w *Worker) reChanKeepEvent() {
	for _, v := range w.keepEvent {
		w.Events <- v
	}
	w.keepEvent = w.keepEvent[len(w.keepEvent):]
}

func (w *Worker) WaitEvent(dur time.Duration, expert ExpectEvent) (Event, error) {
	timeout := time.After(dur)

	for {
		select {
		case event := <-w.Events:
			if !expert.Exist(event.Name) {
				w.keepEvent = append(w.keepEvent, event)
				continue
			}
			fmt.Printf("evented %d\n", len(w.Events))
			w.reChanKeepEvent()
			return event, nil
		case <-timeout:
			fmt.Printf("timeout\n")
			w.reChanKeepEvent()
			return Event{}, errors.New("timeout")
		case <-w.Ctx.Done():
			fmt.Printf("context done\n")
			w.reChanKeepEvent()
			return Event{}, w.Ctx.Err()
		}
	}
}

func (w *Worker) Context() context.Context {
	return w.Ctx
}

type Workers map[string]*Work

var (
	errNotFound     = errors.New("Not Found")
	errCantUpdate   = errors.New("Cant update non-exiting word")
	errWorkerExists = errors.New("That worker already exists")
)

func (w Workers) search(id string) (*Work, error) {
	Work, exist := w[id]
	if exist {
		return Work, nil
	}
	return nil, errNotFound
}

func (w Workers) add(id string, worker *Work) error {
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
