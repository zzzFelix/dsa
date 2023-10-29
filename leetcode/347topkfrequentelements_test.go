package leetcode

import "testing"

func TestTopKFrequent(t *testing.T) {
	t.Run("should return second element", func(t *testing.T) {
		// given
		nums := []int{1, 1, 1, 2, 2, 3}
		expected := []int{1, 2}
		k := 2

		// when
		actual := topKFrequent(nums, k)

		// then
		if !isSlicesEqual(expected, actual) {
			t.Fatalf("did not return expected result. expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("should return only element", func(t *testing.T) {
		// given
		nums := []int{1}
		expected := []int{1}
		k := 1

		// when
		actual := topKFrequent(nums, k)

		// then
		if !isSlicesEqual(expected, actual) {
			t.Fatalf("did not return expected result. expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("should return 7 elements", func(t *testing.T) {
		// given
		nums := []int{5, 1, -1, -8, -7, 8, -5, 0, 1, 10, 8, 0, -4, 3, -1, -1, 4, -5, 4, -3, 0, 2, 2, 2, 4, -2, -4, 8, -7, -7, 2, -8, 0, -8, 10, 8, -8, -2, -9, 4, -7, 6, 6, -1, 4, 2, 8, -3, 5, -9, -3, 6, -8, -5, 5, 10, 2, -5, -1, -5, 1, -3, 7, 0, 8, -2, -3, -1, -5, 4, 7, -9, 0, 2, 10, 4, 4, -4, -1, -1, 6, -8, -9, -1, 9, -9, 3, 5, 1, 6, -1, -2, 4, 2, 4, -6, 4, 4, 5, -5}
		expected := [][]int{
			{4, -1, 2, -5, -8, 0, 8},
			{4, -1, 2, -5, -8, 8, 0},
			{4, -1, 2, -5, 8, -8, 0},
			{4, -1, 2, -5, 8, 0, -8},
			{4, -1, 2, -5, 0, 8, -8},
			{4, -1, 2, -5, 0, -8, 8},
		}

		k := 7

		// when
		actual := topKFrequent(nums, k)

		// then
		foundMatch := false
		for _, val := range expected {
			if isSlicesEqual(val, actual) {
				foundMatch = true
			}
		}
		if !foundMatch {
			t.Fatalf("did not return expected result. expected: %v, actual: %v", expected, actual)
		}
	})
}

func isSlicesEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
