package framework

type Listener interface {
	Handle(e Event) error
}

type ListenerFunc func(e Event) error

func (fn ListenerFunc) Handle(e Event) error {
	return fn(e)
}

// import (
// 	"errors"
// 	"fmt"
// 	"time"
// )

// type Listener interface {
// 	Handle(evt Event) error
// }

// func (w *Worker) reChanKeepEvent() {
// 	for _, v := range w.keepEvent {
// 		w.Events <- v
// 	}
// 	w.keepEvent = w.keepEvent[len(w.keepEvent):]
// }

// func (w *Worker) WaitEvent(dur time.Duration, expert ExpectEvent) (Event, error) {
// 	timeout := time.After(dur)

// 	for {
// 		select {
// 		case event := <-w.Events:
// 			if !expert.Exist(event.Name) {
// 				w.keepEvent = append(w.keepEvent, event)
// 				continue
// 			}
// 			fmt.Printf("evented %d\n", len(w.Events))
// 			w.reChanKeepEvent()
// 			return event, nil
// 		case <-timeout:
// 			fmt.Printf("timeout\n")
// 			w.reChanKeepEvent()
// 			return Event{}, errors.New("timeout")
// 		case <-w.Ctx.Done():
// 			fmt.Printf("context done\n")
// 			w.reChanKeepEvent()
// 			return Event{}, w.Ctx.Err()
// 		}
// 	}
// }
