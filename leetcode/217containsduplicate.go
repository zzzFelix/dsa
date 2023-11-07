package leetcode

import "dsa/core/collection"

func containsDuplicate(nums []int) bool {
	set := make(collection.Set[int])
	for _, val := range nums {
		if set.Contains(val) {
			return true
		}
		set.Insert(val)
	}
	return false
}
