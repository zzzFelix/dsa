package collection

import "errors"

type maxHeap[T comparable] struct {
	items []heapItem[T]
}

type heapItem[T comparable] struct {
	object   T
	priority int
}

type MaxHeap[T comparable] maxHeap[T]

func NewMaxHeap[T comparable](items []T, priorities []int) (*MaxHeap[T], error) {
	size := len(items)
	if size != len(priorities) {
		return nil, errors.New("length of items is not equal to length of priorities")
	}
	mHeap := MaxHeap[T]{items: []heapItem[T]{}}
	for i := range items {
		mHeap.Insert(items[i], priorities[i])
	}
	return &mHeap, nil
}

func (h *MaxHeap[T]) Insert(element T, priority int) {
	mHeapElem := heapItem[T]{
		object:   element,
		priority: priority,
	}
	h.items = append(h.items, mHeapElem)
	h.siftUp()
}

func (h *MaxHeap[T]) Peek() T {
	return h.items[0].object
}

func (h *MaxHeap[T]) Delete() T {
	size := len(h.items)
	removedElement := h.items[0]
	h.items[0] = h.items[size-1]
	h.items = h.items[:size-1]
	h.siftDown(0)
	return removedElement.object
}

func (h *MaxHeap[T]) Size() int {
	return len(h.items)
}

func (h *MaxHeap[T]) swap(i, j int) {
	tmp := h.items[i]
	h.items[i] = h.items[j]
	h.items[j] = tmp
}

func (h *MaxHeap[T]) siftUp() {
	i := len(h.items) - 1
	parent := (i - 1) / 2

	for parent >= 0 && (h.items[i].priority > h.items[parent].priority) {
		h.swap(i, parent)
		i = parent
		parent = (i - 1) / 2
	}
}

func (h *MaxHeap[T]) siftDown(startIndex int) {
	i := startIndex
	leftChild := (2 * i) + 1
	rightChild := (2 * i) + 2

	for (leftChild < len(h.items) && h.items[leftChild].priority > h.items[i].priority) || (rightChild < len(h.items) && h.items[rightChild].priority > h.items[i].priority) {
		if rightChild >= len(h.items) || (h.items[leftChild].priority > h.items[rightChild].priority) {
			h.swap(i, leftChild)
			i = leftChild
		} else {
			h.swap(i, rightChild)
			i = rightChild
		}
		leftChild = (2 * i) + 1
		rightChild = (2 * i) + 2
	}
}
