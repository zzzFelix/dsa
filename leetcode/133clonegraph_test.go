package leetcode

import (
	"dsa/core/common"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		// given
		node4 := common.TreeNode[int]{Value: 4, Children: []common.TreeNode[int]{}}
		node2 := common.TreeNode[int]{Value: 2, Children: []common.TreeNode[int]{}}
		node3 := common.TreeNode[int]{Value: 3, Children: []common.TreeNode[int]{node2, node4}}
		node1 := common.TreeNode[int]{Value: 1, Children: []common.TreeNode[int]{node2, node4}}
		node2.Children = append(node2.Children, node1)
		node2.Children = append(node2.Children, node3)
		node4.Children = append(node4.Children, node1)
		node4.Children = append(node4.Children, node3)

		// when
		actual := cloneGraph(&node1)

		// then
		if node1.Value != actual.Value {
			t.Fatal("Graphs not equal")
		}
		for idx := range node1.Children {
			if node1.Children[idx].Value != actual.Children[idx].Value {
				t.Fatal("Graphs not equal")
			}
			if &node1.Children[idx] == &actual.Children[idx] {
				t.Fatal("Addresses are the same")
			}
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		// given
		node := common.TreeNode[int]{Value: 1, Children: []common.TreeNode[int]{}}

		// when
		actual := cloneGraph(&node)

		// then
		if node.Value != actual.Value {
			t.Fatal("Graphs not equal")
		}
		if &node == actual {
			t.Fatal("Addresses are the same")
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		// when
		actual := cloneGraph(nil)

		// then
		if actual != nil {
			t.Fatal("Graphs not equal")
		}
	})
}
