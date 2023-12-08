package math

import "testing"

func TestLCM(t *testing.T) {
	t.Run("should return LCM of two numbers", func(t *testing.T) {
		// given
		nums := []int{4, 6}
		expected := 12

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

	t.Run("should return LCM of three numbers", func(t *testing.T) {
		// given
		nums := []int{5, 8, 10}
		expected := 40

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

	t.Run("should return error when less than two numbers are provided", func(t *testing.T) {
		// given
		nums := []int{7}
		expected := -1

		// when
		actual, err := LCM(nums...)

		// then
		if err == nil {
			t.Fatal("expected an error, but got nil")
		}
		if expected != actual {
			t.Fatalf("nope! expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("should return LCM of negative numbers", func(t *testing.T) {
		// given
		nums := []int{-3, -5, -9}
		expected := -15

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

	t.Run("should return 0 when one of the numbers is 0", func(t *testing.T) {
		// given
		nums := []int{2, 0, 6}
		expected := 0

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

	t.Run("should return LCM of large numbers", func(t *testing.T) {
		// given
		nums := []int{999999, 888888, 777777}
		expected := 7999992

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
