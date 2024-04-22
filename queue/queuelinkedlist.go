package queue

// node will store the value and the next node as well
type node[T any] struct {
	Data T
	Next *node[T]
}

// LinkedListQueue structure tells us what our head is, what tail should be, and the length of the list
type LinkedListQueue[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func NewLinkedListQueue[T any]() Queue[T] {
	return &LinkedListQueue[T]{}
}

// Enqueue will add a new value into the queue
func (q *LinkedListQueue[T]) Enqueue(n T) error {
	newNode := &node[T]{Data: n, Next: nil} // create and initialize new Node

	if q.tail != nil {
		q.tail.Next = newNode
	}

	q.tail = newNode

	if q.head == nil {
		q.head = newNode
	}
	q.length++

	return nil
}

// Dequeue will remove the first value from the queue (First In First Out)
func (q *LinkedListQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var empty T                 // create a zero value for T
		return empty, ErrQueueEmpty // if is empty return zero value and false
	}
	data := q.head.Data

	q.head = q.head.Next

	if q.head == nil {
		q.tail = nil
	}

	q.length--
	return data, nil
}

// IsEmpty checks if our list is empty or not
func (q *LinkedListQueue[T]) IsEmpty() bool {
	return q.length == 0
}

// Len returns the length of the queue
func (q *LinkedListQueue[T]) Len() int {
	return q.length
}

// FrontQueue returns the front data
func (q *LinkedListQueue[T]) FrontQueue() (T, error) {
	if q.IsEmpty() {
		var empty T
		return empty, ErrQueueEmpty // return zero value and false if empty
	}
	return q.head.Data, nil
}

// BackQueue returns the back data
func (q *LinkedListQueue[T]) BackQueue() (T, error) {
	if q.IsEmpty() {
		var empty T
		return empty, ErrQueueEmpty // return zero value and false if empty
	}
	return q.tail.Data, nil
}
