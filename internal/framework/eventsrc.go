package framework

type Event interface {
	Name() string
}

type EventSrc struct {
	events []Event

	Handler func(Event)
}

func (es *EventSrc) Dispatch(ev Event) {
	es.Handler(ev)
	es.events = append(es.events, ev)
}

func (es *EventSrc) Clear() {
	es.events = []Event{}
}
