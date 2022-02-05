package main

import (
	. "github.com/Gebes/there/v2"
	"log"
)

func main() {
	router := NewRouter()

	err := router.Listen(8080)
	if err != nil {
		log.Fatalln(err)
	}
}
