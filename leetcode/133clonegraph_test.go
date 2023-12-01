package leetcode

import (
	"testing"

	"github.com/zzzFelix/dsa/core/common"
)

func TestCloneGraph(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		// given
		node4 := common.Node[int]{Value: 4, Neighbors: []common.Node[int]{}}
		node2 := common.Node[int]{Value: 2, Neighbors: []common.Node[int]{}}
		node3 := common.Node[int]{Value: 3, Neighbors: []common.Node[int]{node2, node4}}
		node1 := common.Node[int]{Value: 1, Neighbors: []common.Node[int]{node2, node4}}
		node2.Neighbors = append(node2.Neighbors, node1)
		node2.Neighbors = append(node2.Neighbors, node3)
		node4.Neighbors = append(node4.Neighbors, node1)
		node4.Neighbors = append(node4.Neighbors, node3)

		// when
		actual := cloneGraph(&node1)

		// then
		if node1.Value != actual.Value {
			t.Fatal("Graphs not equal")
		}
		for idx := range node1.Neighbors {
			if node1.Neighbors[idx].Value != actual.Neighbors[idx].Value {
				t.Fatal("Graphs not equal")
			}
			if &node1.Neighbors[idx] == &actual.Neighbors[idx] {
				t.Fatal("Addresses are the same")
			}
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		// given
		node := common.Node[int]{Value: 1, Neighbors: []common.Node[int]{}}

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
