package leetcode

import "dsa/core/common"

func maxDepth(root *common.BTreeNode[int]) int {
	return maxDepthHelper(root, 0)
}

func maxDepthHelper(root *common.BTreeNode[int], count int) int {
	if root == nil {
		return count
	} else if root.Left == nil && root.Right == nil {
		return count + 1
	} else {
		leftCount := maxDepthHelper(root.Left, count+1)
		rightCount := maxDepthHelper(root.Right, count+1)
		return max(leftCount, rightCount)
	}
}
