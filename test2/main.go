package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

const MaxEvent = 30

type Result struct {
	eventType uint16
	err       error
	time      time.Time
}

type Handler struct {
	id       uuid.UUID
	timer    *time.Timer
	bus      chan uint16
	expectID string
	close    chan struct{}
}

func NewHandler() *Handler {
	uuid, _ := uuid.NewV4()

	return &Handler{
		id:    uuid,
		bus:   make(chan uint16, MaxEvent),
		close: make(chan struct{}, 1), // Buffered to avoid locking up the event feed
	}
}

func (h *Handler) ID() uuid.UUID              { return h.id }
func (h *Handler) SetTimer(dur time.Duration) { h.timer = time.NewTimer(dur) }
func (h *Handler) ExpectID() string           { return h.expectID }
func (h *Handler) SetExpectID(id string)      { h.expectID = id }

func (h *Handler) Close() {
	close(h.bus)
	if h.close != nil {
		close(h.close)
	}
	h.bus = nil
	h.close = nil
}

func (h *Handler) Wait(ctx context.Context, callID string) <-chan Result {
	eventStream := make(chan Result)
	go func() {
		defer close(eventStream)
		defer h.Close()

		// set default timer
		if h.timer == nil {
			h.timer = time.NewTimer(time.Second * 20)
		}
		// waiting event
		for {
			select {
			case event := <-h.bus:
				fmt.Println("bus input", event)
				eventStream <- Result{eventType: event}
			case <-h.timer.C:
				fmt.Println("event handler timeout")
				eventStream <- Result{err: errors.New("timeout")}
				return
			case <-h.close:
				fmt.Println("event close")
				eventStream <- Result{err: nil}
				return
			case <-ctx.Done():
				fmt.Println("ctx done")
				return
			}
		}
	}()
	return eventStream
}

func main() {
// 	eh := NewHandler()

// 	go func() {
// 		select {
// 		case <-time.After(time.Second * 4):
// 			eh.bus <- 11
// 		}
// 	}()

// 	go func() {
// 		select {
// 		case <-time.After(time.Second * 8):
// 			eh.bus <- 13
// 		}
// 	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

// 	stream := eh.Wait(ctx, "12345")
// loop:
// 	select {
// 	case result := <-stream:
// 		if result.err != nil {
// 			fmt.Println("error is not null", result.err)
// 			return
// 		}
// 		fmt.Println("no error event type:", result)
// 		goto loop
// 	case <-ctx.Done():
// 		fmt.Println("context error")
// 		return
// 	}

	select {
	case <-ctx.Done():
		return
	default:
		fmt.Println("hhhh")
	}
}
