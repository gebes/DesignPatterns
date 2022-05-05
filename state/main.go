package main

import "time"

func main() {
	spaceship := NewSpaceship("Starship Enterprise")
	
	spaceship.Land()

	go func() {
		time.Sleep(time.Second)
		spaceship.Land()
	}()
	spaceship.Launch()
	spaceship.Land()
}
