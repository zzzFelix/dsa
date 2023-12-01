package leetcode

import (
	"testing"

	"github.com/zzzFelix/dsa/core/common"
)

func TestMaxDepth(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		// given
		grandchildLeft := common.BTreeNode[int]{
			Value: 15,
		}
		grandchildRight := common.BTreeNode[int]{
			Value: 7,
		}
		childLeft := common.BTreeNode[int]{
			Value: 9,
		}
		childRight := common.BTreeNode[int]{
			Value: 20,
			Left:  &grandchildLeft,
			Right: &grandchildRight,
		}
		root := common.BTreeNode[int]{
			Value: 3,
			Left:  &childLeft,
			Right: &childRight,
		}
		expected := 3

		// when
		actual := maxDepth(&root)

		// then
		if expected != actual {
			t.Fatalf("test failed. expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		// given
		childRight := common.BTreeNode[int]{
			Value: 2,
		}
		root := common.BTreeNode[int]{
			Value: 1,
			Right: &childRight,
		}
		expected := 2

		// when
		actual := maxDepth(&root)

		// then
		if expected != actual {
			t.Fatalf("test failed. expected: %d, actual: %d", expected, actual)
		}
	})
}
