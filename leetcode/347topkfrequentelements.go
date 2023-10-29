package leetcode

import (
	"dsa/core/collection"
)

func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, val := range nums {
		frequency[val] = frequency[val] + 1
	}

	heap, _ := collection.NewMaxHeap[int]([]int{}, []int{})
	for k, v := range frequency {
		heap.Insert(k, v)
	}

	topK := []int{}
	for i := 0; i < k; i++ {
		topK = append(topK, heap.Delete())
	}

	return topK
}
