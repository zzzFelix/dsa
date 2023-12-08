package math

import "testing"

func TestGCD(t *testing.T) {
	t.Run("should_return_LCM_of_negative_numbers", func(t *testing.T) {
		// given
		nums := []int{-3, -5, -9}
		expected := -45

		// when
		actual := LCM(nums[0], nums[1], nums[2:]...)

		// then
		if expected != actual {
			t.Fatalf("nope! expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("should_return_LCM_of_large_numbers", func(t *testing.T) {
		// given
		nums := []int{999999, 888888, 777777}
		expected := 55999944

		// when
		actual := LCM(nums[0], nums[1], nums[2:]...)

		// then
		if expected != actual {
			t.Fatalf("nope! expected: %d, actual: %d", expected, actual)
		}
	})
}
