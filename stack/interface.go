package stack

import "errors"

type Stack[T any] interface {
	Push(key T) error
	Top() (T, error)
	Pop() (T, error)
	Len() int
	IsEmpty() bool
}

var (
	ErrStackEmpty = errors.New("stack empty")
	ErrStackFull  = errors.New("stack full")
)
