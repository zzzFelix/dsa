package search

import (
	"dsa/core/compare"
	"errors"
)

type BinarySearch[T any] struct {
	SortedItems []T
	Comparator  compare.Comparator[T]
}

func (b *BinarySearch[T]) Search(target T) (int, error) {
	return b.binarySearch(target, 0, len(b.SortedItems))
}

func (b *BinarySearch[T]) binarySearch(target T, left int, right int) (int, error) {
	totalIdxs := right - left
	middleIndex := totalIdxs/2 + left
	if totalIdxs < 0 || middleIndex < 0 || middleIndex > len(b.SortedItems)-1 {
		return -1, errors.New("target item not in sorted array")
	}
	middleItem := b.SortedItems[middleIndex]
	compare := b.Comparator.Compare(middleItem, target)

	if compare == 0 {
		return middleIndex, nil
	}
	if compare < 0 {
		return b.binarySearch(target, middleIndex+1, right)
	} else {
		return b.binarySearch(target, left, middleIndex-1)
	}
}
