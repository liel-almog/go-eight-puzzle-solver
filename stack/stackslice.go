package stack

type SliceStack[T any] struct {
	keys []T
}

func NewSliceStack[T any]() Stack[T] {
	return &SliceStack[T]{}
}

func (stack *SliceStack[T]) Push(key T) error {
	stack.keys = append(stack.keys, key)

	return nil
}

func (stack *SliceStack[T]) Top() (T, error) {
	var x T
	if !stack.IsEmpty() {
		x = stack.keys[len(stack.keys)-1]
		return x, nil
	}

	return x, ErrStackEmpty
}

func (stack *SliceStack[T]) Pop() (T, error) {
	var x T
	if !stack.IsEmpty() {
		x, stack.keys = stack.keys[len(stack.keys)-1], stack.keys[:len(stack.keys)-1]
		return x, nil
	}

	return x, ErrStackEmpty
}

func (stack *SliceStack[T]) IsEmpty() bool {
	return len(stack.keys) == 0
}

func (stack *SliceStack[T]) Len() int {
	return len(stack.keys)
}
