package collection

import "testing"

func TestLinkedList(t *testing.T) {
	t.Run("should create one-element list with no next node", func(t *testing.T) {
		// given
		ll := LinkedList[int]{}
		expVal := 123

		// when
		node := ll.CreateNode(expVal)

		// then
		if node.Next != nil || node.Value != expVal {
			t.Fatal("node should not have a next node")
		}
	})

	t.Run("should create circular two-element list", func(t *testing.T) {
		// given
		ll := LinkedList[struct{}]{}

		// when
		node1 := ll.CreateNode(struct{}{})
		node2 := ll.CreateNode(struct{}{})
		node1.Next = &node2
		node2.Next = &node1

		// then
		if node1.Next != &node2 || node2.Next != &node1 {
			t.Fatal("linked list isn't circular")
		}
	})

	t.Run("should access next node's value", func(t *testing.T) {
		// given
		ll := LinkedList[int]{}

		// when
		node1 := ll.CreateNode(1)
		node2 := ll.CreateNode(2)
		node1.Next = &node2

		// then
		if node1.Next.Value != 2 {
			t.Fatal("couldn't access second node's value")
		}
	})
}
