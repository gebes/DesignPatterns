package observer

type Observer[T any] func(event T)

type Source[T any] interface {
	AddObserver(observer Observer[T])
	NotifyObservers(event T)
}
