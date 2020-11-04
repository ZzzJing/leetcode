/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package tree

func isUnivalTree(root *TreeNode) bool {
	val := root.Val
	return traverse(root, val)
}

func traverse(node *TreeNode, value int) bool {
	if node.Val != value {
		return false
	}

	rsLeft := true
	rsRight := true
	if node.Left != nil {
		rsLeft = traverse(node.Left, value)
	}

	if node.Right != nil {
		rsRight = traverse(node.Right, value)
	}

	if rsLeft && rsRight {
		return true
	} else {
		return false
	}
}
