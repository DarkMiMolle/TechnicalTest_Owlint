package backend

import (
	"bytes"
	"net/http"
)

type Client struct {
	Url string

	RetryPolicy RetryPolicy
}

type InComing[T any] struct {
	inComingElem chan T
	elem         *T
}
type WillSet[T any] struct {
	set_func func(T)
}

func CreateInComingValueOf[T any]() (InComing[T], WillSet[T]) {
	inc := InComing[T]{inComingElem: make(chan T), elem: nil}
	return inc, WillSet[T]{func(value T) {
		inc.inComingElem <- value
	}}
}
func (inc *WillSet[T]) Set(value T) {
	inc.set_func(value)
	inc.set_func = func(T) {
		panic("Value already set")
	}
}
func (inc *InComing[T]) Get() T {
	if inc.elem != nil {
		return *inc.elem
	}
	elem := <-inc.inComingElem
	inc.elem = &elem
	close(inc.inComingElem)
	return elem
}

func (c Client) Post(contentType string, body []byte) (InComing[*http.Response], InComing[error]) {
	inComingResponse, futureSettingResponse := CreateInComingValueOf[*http.Response]()
	inComingError, futureSettingError := CreateInComingValueOf[error]()

	c.RetryPolicy.RunPolicy(func() (*http.Response, error) { return http.Post(c.Url, contentType, bytes.NewBuffer(body)) }, futureSettingResponse, futureSettingError)

	return inComingResponse, inComingError
}
