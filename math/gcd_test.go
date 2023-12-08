package math

import "testing"

func TestGCD(t *testing.T) {
	t.Run("should_return_LCM_of_negative_numbers", func(t *testing.T) {
		// given
		nums := []int{-3, -5, -9}
		expected := 45

		// when
		actual, err := LCM(nums...)

		// then
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if expected != actual {
			t.Fatalf("nope! expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("should_return_LCM_of_large_numbers", func(t *testing.T) {
		// given
		nums := []int{999999, 888888, 777777}
		expected := 617283950617283950

		// when
		actual, err := LCM(nums...)

		// then
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if expected != actual {
			t.Fatalf("nope! expected: %d, actual: %d", expected, actual)
		}
	})
}
