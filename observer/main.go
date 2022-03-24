package main

import (
	observer2 "DesignPatterns/observer"
	"log"
)

func main() {

	var observer observer2.Source[observer2.Event] = &observer2.EventSource{}

	observer.AddObserver(func(event observer2.Event) {
		log.Println("Received a \"" + event.Name + "\" event")
	})

	observer.NotifyObservers(observer2.Event{
		Name:        "Something",
		Description: "It seems like that somethinig happened",
	})

}
