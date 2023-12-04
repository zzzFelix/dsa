package collection

type LinkedListNode[T any] struct {
	Value T
	Next  *LinkedListNode[T]
}

type LinkedList[T any] struct{}

func (l *LinkedList[T]) CreateNode(obj T) LinkedListNode[T] {
	node := LinkedListNode[T]{
		Value: obj,
		Next:  nil,
	}
	return node
}
