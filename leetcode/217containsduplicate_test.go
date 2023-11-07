package leetcode

import "testing"

func TestContainsDuplicate(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		// given
		input := []int{1, 2, 3, 1}
		expected := true

		// when
		actual := containsDuplicate(input)

		// then
		if expected != actual {
			t.Fatalf("did not return correct bool. expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		// given
		input := []int{1, 2, 3, 4}
		expected := false

		// when
		actual := containsDuplicate(input)

		// then
		if expected != actual {
			t.Fatalf("did not return correct bool. expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		// given
		input := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
		expected := true

		// when
		actual := containsDuplicate(input)

		// then
		if expected != actual {
			t.Fatalf("did not return correct bool. expected: %v, actual: %v", expected, actual)
		}
	})
}
