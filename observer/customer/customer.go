package main

import "fmt"

type Customer struct {
	Id string
}

func (c *Customer) Update(itemName string) {
	fmt.Printf("Sending email to observer %s for item %s\n", c.Id, itemName)
}

func (c *Customer) GetID() string {
	return c.Id
}
