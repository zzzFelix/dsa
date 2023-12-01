package leetcode

import "github.com/zzzFelix/dsa/core/collection"

/**
 * https://leetcode.com/problems/reverse-linked-list/
 *
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Value int
 *     Next *collection.Node[int]
 * }
 */
func reverseList(head *collection.Node[int]) *collection.Node[int] {
	return reverseListHelper(head, nil)
}

func reverseListHelper(head *collection.Node[int], prev *collection.Node[int]) *collection.Node[int] {
	if head == nil {
		return prev
	}
	tmp := head.Next
	head.Next = prev
	return reverseListHelper(tmp, head)
}
