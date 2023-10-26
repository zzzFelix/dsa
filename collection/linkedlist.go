package collection

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct{}

func (l *LinkedList[T]) CreateNode(obj T) Node[T] {
	node := Node[T]{
		value: obj,
		next:  nil,
	}
	return node
}
