package main

import (
	"log"
)

func main() {

	var observer Source[Event] = &EventSource{}

	observer.Subscribe(func(event Event) {
		log.Println("Received a \"" + event.Name + "\" event")
	})

	observer.NotifyObservers(Event{
		Name:        "Something",
		Description: "It seems like that somethinig happened",
	})

}
