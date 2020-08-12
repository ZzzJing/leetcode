/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package tree

var preNode *TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	preNode = &TreeNode{}
	head := preNode
	visit(root)
	return head.Right
}

func visit(root *TreeNode) {
	if root == nil {
		return
	}

	visit(root.Left)
	root.Left = nil
	preNode.Right = root
	preNode = root
	visit(root.Right)
}
