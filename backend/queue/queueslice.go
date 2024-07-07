package queue

type SliceQueue[T any] []T

func NewSliceQueue[T any]() Queue[T] {
	return &SliceQueue[T]{}
}

func (q *SliceQueue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *SliceQueue[T]) Dequeue() (T, error) {
	queue := *q

	if q.IsEmpty() {
		var empty T
		return empty, ErrQueueEmpty
	}

	v := queue[0]
	*q = queue[1:]
	return v, nil
}

func (q *SliceQueue[T]) Enqueue(value T) error {
	*q = append(*q, value)
	return nil
}

func (q *SliceQueue[T]) Len() int {
	return len(*q)
}

func (q *SliceQueue[T]) FrontQueue() (T, error) {
	if q.IsEmpty() {
		var empty T
		return empty, nil
	}

	v := (*q)[0]

	return v, nil
}

func (q *SliceQueue[T]) BackQueue() (T, error) {
	if q.IsEmpty() {
		var empty T
		return empty, nil
	}

	v := (*q)[q.Len()-1]

	return v, nil
}
