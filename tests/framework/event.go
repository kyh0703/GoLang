package framework

import (
	"context"
	"errors"
	"time"
)

const (
	MAX_EVENTS     = 10
	EVENT_CHAN_KEY = "recvChan"
)

var (
	errDiffChanType    = errors.New("did not found Chan Type")
	errContextNotFound = errors.New("did not found Context")
	errChanNotFound    = errors.New("did not found Chan")
	errTimeout         = errors.New("timeout")
)

type EventType uint8

type Event struct {
	Id   string
	Type EventType
	Data interface{}
}

type ExpectEvent struct {
	Id   string
	Type EventType
}

type Listener interface {
	Handle(e Event) error
}

type ListenerFunc func(e Event) error

func getRecvChan(ctx context.Context) (chan Event, error) {
	v := ctx.Value(EVENT_CHAN_KEY)
	if v == nil {
		return nil, errChanNotFound
	}

	ch, ok := v.(chan Event)
	if !ok {
		return nil, errDiffChanType
	}
	return ch, nil
}

func reRegisEvent(ch chan Event, events []Event) {
	for _, v := range events {
		ch <- v
	}
}

func WaitEvent(ctx context.Context, timeout time.Duration,
	listener ListenerFunc, expects ...ExpectEvent) error {
	if ctx == nil {
		return errContextNotFound
	}

	recvChan, err := getRecvChan(ctx)
	if err != nil {
		return err
	}

	tempSaveEvents := make([]Event, 0, MAX_EVENTS)

	for {
		select {
		case event := <-recvChan:
			var isFind bool = false
			for _, v := range expects {
				if event.Type == v.Type && event.Id == v.Id {
					isFind = true
					break
				}
			}

			if !isFind {
				tempSaveEvents = append(tempSaveEvents, event)
				continue
			}

			reRegisEvent(recvChan, tempSaveEvents)
			return listener(event)
		case <-time.After(timeout):
			reRegisEvent(recvChan, tempSaveEvents)
			return errTimeout
		case <-ctx.Done():
			reRegisEvent(recvChan, tempSaveEvents)
			return ctx.Err()
		}
	}
}
