package main

type Observer interface {
	Update(string)
	GetID() string
}
