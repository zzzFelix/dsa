package collection

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct{}

func (l *LinkedList[T]) CreateNode(obj T) Node[T] {
	node := Node[T]{
		Value: obj,
		Next:  nil,
	}
	return node
}
