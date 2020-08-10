/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package tree

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	node := &TreeNode{
		Val: t1.Val + t2.Val,
	}

	if t1.Left != nil || t2.Left != nil {
		node.Left = mergeTrees(t1.Left, t2.Left)
	}

	if t1.Right != nil || t2.Right != nil {
		node.Right = mergeTrees(t1.Right, t2.Right)
	}

	return node
}
