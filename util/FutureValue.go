package util

/*
InComing represents an in coming value that is currently computed (in another goroutine).

Use CreateInComingValueOf to create an InComing value.

Note: it should not be possible to get an InComing value without having a goroutine.
*/
type InComing[T any] struct {
	inComingElem chan T
	elem         *T
}

// Get wait for the value and/or return it when she is ready
func (inc *InComing[T]) Get() T {
	if inc.elem != nil {
		return *inc.elem
	}
	elem := <-inc.inComingElem
	inc.elem = &elem
	close(inc.inComingElem)
	return elem
}

/*
WillSet works on the other side of the InComing value. It allows to set the value of the InComing value.
Once a value set, it won't be possible to set another value and trying to do so will panic

Use CreateInComingValueOf to create a WillSet value.

Note: It should not be possible to get a WillSet value without having a goroutine.
*/
type WillSet[T any] struct {
	set_func func(T)
}

// Set the value for the InComing value associated to. Calling it more than once will panic.
func (inc *WillSet[T]) Set(value T) {
	inc.set_func(value)
	inc.set_func = func(T) {
		panic("Value already set")
	}
}

// CreateInComingValueOf create a duo InComing / WillSet value.
//
//Note: CreateInComingValueOf should only be called near a goroutine run
func CreateInComingValueOf[T any]() (InComing[T], WillSet[T]) {
	inc := InComing[T]{inComingElem: make(chan T), elem: nil}
	return inc, WillSet[T]{func(value T) {
		inc.inComingElem <- value
	}}
}
