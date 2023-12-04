package internal

type Node[T comparable] struct {
	Value     T
	Neighbors []Node[T]
}
