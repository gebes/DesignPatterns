package main

type Observer[T any] func(event T)

type Source[T any] interface {
	Subscribe(observer Observer[T])
	NotifyObservers(event T)
}
