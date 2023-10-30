package common

type TreeNode[T comparable] struct {
	Value    T
	Children []TreeNode[T]
}
