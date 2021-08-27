package main

import (
	"fmt"
	"framework"
	"log"
	"strconv"
	"time"
	"transaction"
)

type Event struct {
	Name  string
	Event interface{}
}

type ExpectEvent struct {
	eventName []string
}

func NewExpectEvent() ExpectEvent {
	events := ExpectEvent{
		eventName: make([]string, 0, 10),
	}

	return events
}

func (e ExpectEvent) Add(event string) ExpectEvent {
	e.eventName = append(e.eventName, event)
	return e
}

func (e ExpectEvent) Exist(event string) bool {
	for _, v := range e.eventName {
		if v == event {
			return true
		}
	}
	return false
}

type ListenerFunc func(expect ExpectEvent) error

func WaitEvent(fn ListenerFunc, expect ExpectEvent) bool {
	fn(expect)
	return true
}

func test(expect ExpectEvent) error {
	fmt.Println(expect)
	return nil
}

func main() {
	// WaitEvent(test, NewExpectEvent().Add("test"))
	manager := framework.NewManager(transaction.MakeTransaction)
	for i := 0; i < 100; i++ {
		manager.BeginWorker(strconv.Itoa(i), transaction.Make3PCC)
	}
	time.Sleep(time.Second * 10)
	log.Println("Wait Group")
	manager.EndWorkers()
	manager.PrintWorker()
}
