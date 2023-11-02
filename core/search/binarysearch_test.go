package search

import (
	"dsa/core/compare"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	t.Run("should find middle number of 3", func(t *testing.T) {
		// given
		bSearch := BinarySearch[int]{
			SortedItems: []int{1, 2, 3},
			Comparator:  compare.IntComparator{},
		}

		// when
		index, error := bSearch.Search(2)

		// then
		if error != nil {
			t.Fatalf("Search threw error: %v", error)
		}
		if index != 1 {
			t.Fatalf("Search returned wrong index. Expected: 1, actual: %d", index)
		}
	})

	t.Run("should find left item of 2", func(t *testing.T) {
		// given
		bSearch := BinarySearch[int]{
			SortedItems: []int{1, 2},
			Comparator:  compare.IntComparator{},
		}

		// when
		index, error := bSearch.Search(1)

		// then
		if error != nil {
			t.Fatalf("Search threw error: %v", error)
		}
		if index != 0 {
			t.Fatalf("Search returned wrong index. Expected: 0, actual: %d", index)
		}
	})

	t.Run("should find right item of 2", func(t *testing.T) {
		// given
		bSearch := BinarySearch[int]{
			SortedItems: []int{1, 2},
			Comparator:  compare.IntComparator{},
		}

		// when
		index, error := bSearch.Search(2)

		// then
		if error != nil {
			t.Fatalf("Search threw error: %v", error)
		}
		if index != 1 {
			t.Fatalf("Search returned wrong index. Expected: 1, actual: %d", index)
		}
	})

	t.Run("should find single item", func(t *testing.T) {
		// given
		bSearch := BinarySearch[int]{
			SortedItems: []int{1},
			Comparator:  compare.IntComparator{},
		}

		// when
		index, error := bSearch.Search(1)

		// then
		if error != nil {
			t.Fatalf("Search threw error: %v", error)
		}
		if index != 0 {
			t.Fatalf("Search returned wrong index. Expected: 0, actual: %d", index)
		}
	})

	t.Run("should return error for single item", func(t *testing.T) {
		// given
		bSearch := BinarySearch[int]{
			SortedItems: []int{1},
			Comparator:  compare.IntComparator{},
		}

		// when
		_, error := bSearch.Search(0)

		// then
		if error == nil {
			t.Fatal("Search have thrown error")
		}
	})

	t.Run("should find left-most item of many", func(t *testing.T) {
		// given
		bSearch := BinarySearch[int]{
			SortedItems: []int{1, 2, 3, 4, 5, 6, 99, 102, 4065, 423423, 423424, 423425, 423426},
			Comparator:  compare.IntComparator{},
		}

		// when
		index, error := bSearch.Search(1)

		// then
		if error != nil {
			t.Fatalf("Search threw error: %v", error)
		}
		if index != 0 {
			t.Fatalf("Search returned wrong index. Expected: 0, actual: %d", index)
		}
	})

	t.Run("should find right-most item of many", func(t *testing.T) {
		// given
		bSearch := BinarySearch[int]{
			SortedItems: []int{1, 2, 3, 4, 5, 6, 99, 102, 4065, 423423, 423424, 423425, 423426},
			Comparator:  compare.IntComparator{},
		}

		// when
		index, error := bSearch.Search(423426)

		// then
		if error != nil {
			t.Fatalf("Search threw error: %v", error)
		}
		if index != 12 {
			t.Fatalf("Search returned wrong index. Expected: 12, actual: %d", index)
		}
	})

	t.Run("should find random item of many", func(t *testing.T) {
		// given
		bSearch := BinarySearch[int]{
			SortedItems: []int{1, 2, 3, 4, 5, 6, 99, 102, 4065, 423423, 423424, 423425, 423426},
			Comparator:  compare.IntComparator{},
		}

		// when
		index, error := bSearch.Search(5)

		// then
		if error != nil {
			t.Fatalf("Search threw error: %v", error)
		}
		if index != 4 {
			t.Fatalf("Search returned wrong index. Expected: 4, actual: %d", index)
		}
	})

	t.Run("should return error when item not in list", func(t *testing.T) {
		// given
		bSearch := BinarySearch[int]{
			SortedItems: []int{1, 2, 3, 4, 5, 6, 99, 102, 4065, 423423, 423424, 423425, 423426},
			Comparator:  compare.IntComparator{},
		}

		// when
		_, error := bSearch.Search(8)

		// then
		if error == nil {
			t.Fatal("Search have thrown error")
		}
	})

	t.Run("should return error when list empty", func(t *testing.T) {
		// given
		bSearch := BinarySearch[int]{
			SortedItems: []int{},
			Comparator:  compare.IntComparator{},
		}

		// when
		_, error := bSearch.Search(0)

		// then
		if error == nil {
			t.Fatal("Search have thrown error")
		}
	})
}
