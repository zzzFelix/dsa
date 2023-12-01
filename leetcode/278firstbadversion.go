package leetcode

import (
	bsearch "github.com/zzzFelix/dsa/core/search"
)

func firstBadVersion(n int, badVersion int) int {
	versions := createVersionArray(n)
	bsearch := bsearch.BinarySearch[int]{
		SortedItems: versions,
		Comparator:  BadVersionComparator{},
	}
	i, _ := bsearch.Search(badVersion)
	return i + 1
}

func isBadVersion(n int, badVersion int) bool {
	return n == badVersion
}

func createVersionArray(n int) []int {
	versions := make([]int, n)
	for i := 0; i < n; i++ {
		versions[i] = i + 1
	}
	return versions
}

type BadVersionComparator struct{}

func (c BadVersionComparator) Compare(a, b int) int {
	if isBadVersion(a, b) {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
