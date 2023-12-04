package internal

type BTreeNode[T comparable] struct {
	Value T
	Left  *BTreeNode[T]
	Right *BTreeNode[T]
}
