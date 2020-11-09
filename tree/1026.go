package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var maxDiff int

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	max := root.Val
	min := root.Val
	maxDiff = 0
	genMaxDiff(root, max, min)
	return maxDiff
}

func genMaxDiff(root *TreeNode, max, min int) {
	diffWithMax := absVal(root.Val, max)
	if diffWithMax > maxDiff {
		maxDiff = diffWithMax
	}
	diffWithMin := absVal(root.Val, min)
	if diffWithMin > maxDiff {
		maxDiff = diffWithMin
	}
	if root.Val > max {
		max = root.Val
	}

	if root.Val < min {
		min = root.Val
	}

	if root.Left != nil {
		genMaxDiff(root.Left, max, min)
	}

	if root.Right != nil {
		genMaxDiff(root.Right, max, min)
	}
}

func absVal(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
