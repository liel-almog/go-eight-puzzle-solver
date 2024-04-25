package heap

type SliceHeap[T any] struct {
	heaps    []T
	lessFunc func(a, b T) bool
}

func NewSliceHeap[T any](less func(a, b T) bool) Heap[T] {
	return &SliceHeap[T]{
		lessFunc: less,
	}
}

func (heap *SliceHeap[T]) Push(x T) error {
	heap.heaps = append(heap.heaps, x)
	heap.up(len(heap.heaps) - 1)

	return nil
}

func (heap *SliceHeap[T]) Len() int {
	return len(heap.heaps)
}

func (heap *SliceHeap[T]) Pop() (T, error) {
	if heap.IsEmpty() {
		heap.heaps = nil
		var empty T
		return empty, ErrHeapEmpty
	}

	var x = heap.heaps[0]
	heap.swap(0, len(heap.heaps)-1)
	heap.heaps = heap.heaps[:len(heap.heaps)-1]
	heap.down(0)

	return x, nil
}

func (heap *SliceHeap[T]) IsEmpty() bool {
	return heap.Len() == 0
}

func (heap *SliceHeap[T]) Top() (T, error) {
	if !heap.IsEmpty() {
		return heap.heaps[0], nil
	}

	var empty T
	return empty, ErrHeapEmpty
}

func (heap *SliceHeap[T]) down(parent int) {
	lessIdx := parent
	lChild, rChild := (parent<<1)+1, (parent<<1)+2
	if lChild < len(heap.heaps) && heap.lessFunc(heap.heaps[lChild], heap.heaps[lessIdx]) {
		lessIdx = lChild
	}
	if rChild < len(heap.heaps) && heap.lessFunc(heap.heaps[rChild], heap.heaps[lessIdx]) {
		lessIdx = rChild
	}
	if lessIdx == parent {
		return
	}
	heap.swap(lessIdx, parent)
	heap.down(lessIdx)
}

func (heap *SliceHeap[T]) swap(i, j int) {
	heap.heaps[i], heap.heaps[j] = heap.heaps[j], heap.heaps[i]
}

func (heap *SliceHeap[T]) up(child int) {
	if child <= 0 {
		return
	}
	parent := (child - 1) >> 1
	if !heap.lessFunc(heap.heaps[child], heap.heaps[parent]) {
		return
	}

	heap.swap(child, parent)
	heap.up(parent)
}
