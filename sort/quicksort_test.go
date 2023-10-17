package sort

import (
	"fmt"
	"testing"
)

func TestSortIntegerSlice(t *testing.T) {
	t.Run("should sort slice with negative and positive numbers", func(t *testing.T) {
		// given
		unsorted := []int{99, -2, 3, 4, -88213, 7890, 0, -9, 1, -7777}

		// when
		sorted := quicksort(unsorted)

		// then
		expected := []int{-88213, -7777, -9, -2, 0, 1, 3, 4, 99, 7890}
		if !areSlicesEqual(expected, sorted) {
			t.Fatal(formatDiff(expected, sorted))
		}
	})

	t.Run("should sort slice with only negative numbers", func(t *testing.T) {
		// given
		unsorted := []int{-99, -2, -3, -4, -88213, -7890, -77, -9, -1, -7777}

		// when
		sorted := quicksort(unsorted)

		// then
		expected := []int{-88213, -7890, -7777, -99, -77, -9, -4, -3, -2, -1}
		if !areSlicesEqual(expected, sorted) {
			t.Fatal(formatDiff(expected, sorted))
		}
	})

	t.Run("should sort slice with duplicate numbers", func(t *testing.T) {
		// given
		unsorted := []int{99, 2, 99, 4, 99, 7890, 0, 99, 1, 99}

		// when
		sorted := quicksort(unsorted)

		// then
		expected := []int{0, 1, 2, 4, 99, 99, 99, 99, 99, 7890}
		if !areSlicesEqual(expected, sorted) {
			t.Fatal(formatDiff(expected, sorted))
		}
	})

	t.Run("should sort already sorted slice", func(t *testing.T) {
		// given
		unsorted := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

		// when
		sorted := quicksort(unsorted)

		// then
		expected := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
		if !areSlicesEqual(expected, sorted) {
			t.Fatal(formatDiff(expected, sorted))
		}
	})

	t.Run("should sort reversed slice", func(t *testing.T) {
		// given
		unsorted := []int{19, 17, 15, 13, 11, 9, 7, 5, 3, 1}

		// when
		sorted := quicksort(unsorted)

		// then
		expected := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
		if !areSlicesEqual(expected, sorted) {
			t.Fatal(formatDiff(expected, sorted))
		}
	})

	t.Run("should sort empty slice", func(t *testing.T) {
		// given
		unsorted := []int{}

		// when
		sorted := quicksort(unsorted)

		// then
		expected := []int{}
		if !areSlicesEqual(expected, sorted) {
			t.Fatal(formatDiff(expected, sorted))
		}
	})

	t.Run("should sort slice with 1 item", func(t *testing.T) {
		// given
		unsorted := []int{1}

		// when
		sorted := quicksort(unsorted)

		// then
		expected := []int{1}
		if !areSlicesEqual(expected, sorted) {
			t.Fatal(formatDiff(expected, sorted))
		}
	})

	t.Run("should sort slice with 2 items", func(t *testing.T) {
		// given
		unsorted := []int{2, 1}

		// when
		sorted := quicksort(unsorted)

		// then
		expected := []int{1, 2}
		if !areSlicesEqual(expected, sorted) {
			t.Fatal(formatDiff(expected, sorted))
		}
	})

	t.Run("should sort slice with 3 items", func(t *testing.T) {
		// given
		unsorted := []int{2, 1, 3}

		// when
		sorted := quicksort(unsorted)

		// then
		expected := []int{1, 2, 3}
		if !areSlicesEqual(expected, sorted) {
			t.Fatal(formatDiff(expected, sorted))
		}
	})
}

func areSlicesEqual(a []int, b []int) bool {
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

func formatDiff(expected []int, actual []int) string {
	return fmt.Sprintf("\n[EXPECTED] %v\n[ ACTUAL ] %v", expected, actual)
}
