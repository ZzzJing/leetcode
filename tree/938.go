/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package tree

func rangeSumBST(root *TreeNode, L int, R int) int {
	result := 0
	left := 0
	right := 0

	if root == nil {
		return 0
	}

	if root.Val >= L && root.Val <= R {
		result += root.Val
	}

	if root.Val >= L {
		left = rangeSumBST(root.Left, L, R)
	}

	if root.Val <= R {
		right = rangeSumBST(root.Right, L, R)
	}

	return result + left + right
}
