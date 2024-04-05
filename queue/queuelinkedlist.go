package queue

// Node will store the value and the next node as well
type Node[T any] struct {
	Data T
	Next *Node[T]
}

// Queue structure tells us what our head is, what tail should be, and the length of the list
type Queue[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

// Enqueue will add a new value into the queue
func (q *Queue[T]) Enqueue(n T) {
	newNode := &Node[T]{Data: n} // create and initialize new Node

	if q.tail != nil {
		q.tail.Next = newNode
	}

	q.tail = newNode

	if q.head == nil {
		q.head = newNode
	}
	q.length++
}

// Dequeue will remove the first value from the queue (First In First Out)
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zeroVal T         // create a zero value for T
		return zeroVal, false // if is empty return zero value and false
	}
	data := q.head.Data

	q.head = q.head.Next

	if q.head == nil {
		q.tail = nil
	}

	q.length--
	return data, true
}

// IsEmpty checks if our list is empty or not
func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}

// Len returns the length of the queue
func (q *Queue[T]) Len() int {
	return q.length
}

// FrontQueue returns the front data
func (q *Queue[T]) FrontQueue() (T, bool) {
	if q.IsEmpty() {
		var zeroVal T
		return zeroVal, false // return zero value and false if empty
	}
	return q.head.Data, true
}

// BackQueue returns the back data
func (q *Queue[T]) BackQueue() (T, bool) {
	if q.IsEmpty() {
		var zeroVal T
		return zeroVal, false // return zero value and false if empty
	}
	return q.tail.Data, true
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}
