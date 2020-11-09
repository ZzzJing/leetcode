package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := 0
	changeNode538(root, &sum)
	return root
}

func changeNode538(root *TreeNode, sum *int) {
	if root.Right != nil {
		changeNode538(root.Right, sum)
	}

	root.Val = root.Val + *sum
	*sum = root.Val

	if root.Left != nil {
		changeNode538(root.Left, sum)
	}
}
