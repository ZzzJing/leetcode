package tree

import "math"

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := math.MinInt64
	var maxOne func(node *TreeNode) int
	maxOne = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max124(0, maxOne(node.Left))
		right := max124(0, maxOne(node.Right))
		ans = max124(ans, left+right+node.Val)
		return max124(left, right) + node.Val
	}

	maxOne(root)
	return ans
}

func max124(q, p int) int {
	if q > p {
		return q
	} else {
		return p
	}
}
