package leetcode

import (
	"testing"

	"github.com/zzzFelix/dsa/core/collection"
)

func TestReverseLinkedList(t *testing.T) {
	t.Run("reverse 5 element list", func(t *testing.T) {
		// given
		ll := collection.LinkedList[int]{}
		node1 := ll.CreateNode(1)
		node2 := ll.CreateNode(2)
		node3 := ll.CreateNode(3)
		node4 := ll.CreateNode(4)
		node5 := ll.CreateNode(5)
		node1.Next = &node2
		node2.Next = &node3
		node3.Next = &node4
		node4.Next = &node5

		// when
		actual := reverseList(&node1)

		// then
		expected := []int{5, 4, 3, 2, 1}
		current := actual
		for _, val := range expected {
			if current.Value != val {
				t.Fatal("values aren't equal")
			}
			current = current.Next
		}
	})

	t.Run("reverse 2 element list", func(t *testing.T) {
		// given
		ll := collection.LinkedList[int]{}
		node1 := ll.CreateNode(1)
		node2 := ll.CreateNode(2)
		node1.Next = &node2

		// when
		actual := reverseList(&node1)

		// then
		expected := []int{2, 1}
		current := actual
		for _, val := range expected {
			if current.Value != val {
				t.Fatal("values aren't equal")
			}
			current = current.Next
		}
	})

	t.Run("reverse empty list", func(t *testing.T) {
		// given
		var expected *collection.Node[int] = nil

		// when
		actual := reverseList(expected)

		// then
		if expected != actual {
			t.Fatal("result list is not empty")
		}
	})
}
