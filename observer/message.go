package main

type Event struct {
	Name        string
	Description string
}

type EventSource struct {
	Observers []Observer[Event]
}

func (e *EventSource) AddObserver(observer Observer[Event]) {
	e.Observers = append(e.Observers, observer)
}

func (e *EventSource) NotifyObservers(event Event) {
	for _, observer := range e.Observers {
		observer(event)
	}
}
