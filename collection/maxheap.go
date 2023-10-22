package collection

import (
	"errors"
)

type maxHeap[T comparable] struct {
	items []heapItem[T]
}

type heapItem[T comparable] struct {
	object   T
	priority int
}

type MaxHeap[T comparable] maxHeap[T]

func New[T comparable](items []T, priorities []int) (*MaxHeap[T], error) {
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
	h.siftUp(h.Size() - 1)
}

func (h *MaxHeap[T]) Peek() T {
	return h.items[0].object
}

func (h *MaxHeap[T]) Pop() T {
	size := len(h.items)
	elem := h.items[size-1]
	h.swap(0, size-1)
	h.items = h.items[:size-1]
	h.siftDown(0)
	return elem.object
}

func (h *MaxHeap[T]) Size() int {
	return len(h.items)
}

func (h *MaxHeap[T]) swap(i, j int) {
	tmp := h.items[i]
	h.items[i] = h.items[j]
	h.items[j] = tmp
}

func (h *MaxHeap[T]) siftUp(i int) {
	current := h.items[i].priority
	parentI := (i - 1) / 2
	if parentI < 0 {
		return
	}
	parent := h.items[parentI].priority
	if current > parent {
		h.swap(parentI, i)
		h.siftUp(parentI)
	}
}

func (h *MaxHeap[T]) siftDown(i int) {
	childAI := (2 * i) + 1
	childBI := (2 * i) + 2
	if childAI >= len(h.items) {
		return
	}
	current := h.items[i].priority
	childA := h.items[childAI].priority
	if childBI >= len(h.items) {
		if current < childA {
			h.swap(childAI, i)
			h.siftDown(childAI)
		}
		return
	}
	childB := h.items[childBI].priority
	if childA < childB || current < childB {
		h.swap(childBI, i)
		h.siftDown(childBI)
	} else {
		h.swap(childAI, i)
		h.siftDown(childAI)
	}
}
