package main

import (
	"log"
	"time"
)

type Spaceship struct {
	Name  string
	State EngineState
}

func NewSpaceship(name string) *Spaceship {
	return &Spaceship{Name: name, State: EngineStateOff}
}

func (s *Spaceship) Launch() {
	s.State.Starting()
	log.Println(s.Name + ": Engine is starting...")
	time.Sleep(2 * time.Second)
	s.State.Running()
	log.Println(s.Name + ": Engine is running...")
}

func (s *Spaceship) Land() {
	if s.State == EngineStateOff {
		log.Println(s.Name + ": Engine already off.")
		return
	}
	if s.State == EngineStateStarting {
		log.Println(s.Name + ": Engine is starting. Cannot cancel procedure.")
		return
	}
	s.State.Stop()
	log.Println(s.Name + ": Engine stopped.")
}
