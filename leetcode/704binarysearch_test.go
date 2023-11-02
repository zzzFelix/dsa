package leetcode

import "testing"

func TestBinarySearch(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		// given
		nums := []int{-1, 0, 3, 5, 9, 12}
		target := 9
		expected := 4

		// when
		actual := search(nums, target)

		// then
		if expected != actual {
			t.Fatalf("wrong index. expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		// given
		nums := []int{-1, 0, 3, 5, 9, 12}
		target := 2
		expected := -1

		// when
		actual := search(nums, target)

		// then
		if expected != actual {
			t.Fatalf("wrong index. expected: %d, actual: %d", expected, actual)
		}
	})
}
