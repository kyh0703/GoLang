package framework

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
