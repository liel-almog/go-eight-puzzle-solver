package heap

import "errors"

type Heap[T any] interface {
	Len() int
	Push(key T) error
	Pop() (T, error)
	Top() (T, error)
	IsEmpty() bool
}

var (
	ErrLessFuncRequired = errors.New("less func is necessary")
	ErrHeapEmpty        = errors.New("heap empty")
)
