package framework

import "testing"

type testEvent struct{}

func (*testEvent) Name() string { return "test" }

func TestDispatch(t *testing.T) {
	var raised Event = &testEvent{}
	var handled Event = nil

	src := EventSrc{
		Handler: func(e Event) {
			handled = e
		},
	}

	src.Dispatch(raised)

	if raised != handled {
		t.Error("Raised event not handled")
	}

	if len(src.events) != 1 || src.events[0] != raised {
		t.Error("Raised event not stored")
	}
}

func TestClear(t *testing.T) {
	src := EventSrc{
		events: []Event{
			&testEvent{},
			&testEvent{},
			&testEvent{},
			&testEvent{},
		},
	}

	src.Clear()

	if len(src.events) != 0 {
		t.Error("Events not cleared")
	}
}
