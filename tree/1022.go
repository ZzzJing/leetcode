/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package tree

func sumRootToLeaf(root *TreeNode) int {
	return add(root, 0)
}

func add(root *TreeNode, count int) int {
	if root != nil {
		count = 2*count + root.Val

		if root.Left == nil && root.Right == nil {
			return count
		}

		return add(root.Left, count) + add(root.Right, count)
	}

	return 0
}
