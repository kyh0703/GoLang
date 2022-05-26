package event

import "kafkaevent"

const (
	TestType = iota
)

func Test() {
	kafkaevent.Hello()
}
