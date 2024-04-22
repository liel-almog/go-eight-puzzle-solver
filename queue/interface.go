package queue

import "errors"

type Queue[T any] interface {
	Enqueue(T) error
	Dequeue() (T, error)
	Len() int
	FrontQueue() (T, error)
	BackQueue() (T, error)
	IsEmpty() bool
}

var (
	ErrQueueEmpty = errors.New("queue empty")
	ErrQueueFull  = errors.New("queue full")
)
