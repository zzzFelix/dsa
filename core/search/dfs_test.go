package search

import "testing"

func TestDfs(t *testing.T) {
	t.Run("should find target when given single node", func(t *testing.T) {
		// given
		target := 5
		node := Node[int]{value: target, children: []Node[int]{}}

		// when
		actual, err := node.dfs(target)

		// then
		if actual != target || err != nil {
			t.Fatal("could not find target")
		}
	})

	t.Run("should find target when given multiple nodes", func(t *testing.T) {
		// given
		target := 5
		grandchildA := Node[int]{value: 99, children: []Node[int]{}}
		grandchildB := Node[int]{value: target, children: []Node[int]{}}
		childA := Node[int]{value: 99, children: []Node[int]{}}
		childB := Node[int]{value: 99, children: []Node[int]{grandchildA, grandchildB}}
		node := Node[int]{value: 99, children: []Node[int]{childA, childB}}

		// when
		actual, err := node.dfs(target)

		// then
		if actual != target || err != nil {
			t.Fatal("could not find target")
		}
	})

	t.Run("should return error when target is not in nodes", func(t *testing.T) {
		// given
		grandchildA := Node[int]{value: 99, children: []Node[int]{}}
		grandchildB := Node[int]{value: 99, children: []Node[int]{}}
		childA := Node[int]{value: 99, children: []Node[int]{}}
		childB := Node[int]{value: 99, children: []Node[int]{grandchildA, grandchildB}}
		node := Node[int]{value: 99, children: []Node[int]{childA, childB}}

		// when
		actual, err := node.dfs(5)

		// then
		if err == nil {
			t.Fatalf("did not throw error. actual: %d", actual)
		}
	})
}
