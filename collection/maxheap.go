package collection

import "errors"

type heapElement[T comparable] struct {
	object   T
	priority int
}

type MaxHeap[T comparable] []heapElement[T]

func (h MaxHeap[T]) Insert(element T, priority int) {
	mHeapElem := heapElement[T]{
		object:   element,
		priority: priority,
	}
	h[0] = mHeapElem
}

func (h MaxHeap[T]) Peek() T {
	return h[0].object
}

func (h MaxHeap[T]) siftUp() {

}

func (h MaxHeap[T]) siftDown() {

}

func New[T comparable](elements []T, priorities []int) (MaxHeap[T], error) {
	size := len(elements)
	if size != len(priorities) {
		return nil, errors.New("length of elements is not equal to length of priorities")
	}
	mHeap := make(MaxHeap[T], size)
	for i := range elements {
		mHeap[i] = heapElement[T]{
			object:   elements[i],
			priority: priorities[i],
		}
	}
	return mHeap, nil
}
