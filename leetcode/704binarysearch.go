package leetcode

import (
	"dsa/core/compare"
	dsasearch "dsa/core/search"
)

func search(nums []int, target int) int {
	bs := dsasearch.BinarySearch[int]{
		SortedItems: nums,
		Comparator:  compare.IntComparator{},
	}

	val, _ := bs.Search(target)
	return val
}
