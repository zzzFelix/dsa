package math

import "testing"

func TestLCM(t *testing.T) {
	t.Run("should return LCM of two numbers", func(t *testing.T) {
		// given
		nums := []int{4, 6}
		expected := 12

		// when
		actual := LCM(nums[0], nums[1])

		// then
		if expected != actual {
			t.Fatalf("nope! expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("should return LCM of three numbers", func(t *testing.T) {
		// given
		nums := []int{5, 8, 10}
		expected := 40

		// when
		actual := LCM(nums[0], nums[1], nums...)

		// then
		if expected != actual {
			t.Fatalf("nope! expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("should return LCM of negative numbers", func(t *testing.T) {
		// given
		nums := []int{-3, -5, -9}
		expected := -45

		// when
		actual := LCM(nums[0], nums[1], nums[2])

		// then
		if expected != actual {
			t.Fatalf("nope! expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("should return 0 when one of the numbers is 0", func(t *testing.T) {
		// given
		nums := []int{2, 0, 6}
		expected := 0

		// when
		actual := LCM(nums[0], nums[1], nums[2])

		// then
		if expected != actual {
			t.Fatalf("nope! expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("should return LCM of large numbers", func(t *testing.T) {
		// given
		nums := []int{999999, 888888, 777777}
		expected := 55999944

		// when
		actual := LCM(nums[0], nums[1], nums[2])

		// then
		if expected != actual {
			t.Fatalf("nope! expected: %d, actual: %d", expected, actual)
		}
	})
}
