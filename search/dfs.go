package search

import "errors"

type Node[T comparable] struct {
	value    T
	children []Node[T]
}

func (n *Node[T]) dfs(target T) (T, error) {
	if n.value == target {
		return n.value, nil
	}
	for _, child := range n.children {
		value, err := child.dfs(target)
		if err == nil {
			return value, nil
		}
	}
	return target, errors.New("target is not in nodes")
}
