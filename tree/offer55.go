package tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	return recur(root) != -1
}

func recur(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := recur(root.Left)
	if left == -1 {
		return -1
	}

	right := recur(root.Right)
	if right == -1 {
		return -1
	}

	if math.Abs(float64(left)-float64(right)) <= 1 {
		return int(math.Max(float64(left), float64(right))) + 1
	} else {
		return -1
	}

}
