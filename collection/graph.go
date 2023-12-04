package collection

type GraphNode[T comparable] struct {
	Value     T
	Neighbors []GraphNode[T]
}
