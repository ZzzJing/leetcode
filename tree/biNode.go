package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBiNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	cur := &TreeNode{}
	head := cur
	changeNode(root, cur)
	return head.Right
}

func changeNode(root *TreeNode, cur *TreeNode) *TreeNode {
	if root.Left != nil {
		cur = changeNode(root.Left, cur)
	}

	root.Left = nil
	cur.Right = root
	cur = root

	if root.Right != nil {
		cur = changeNode(root.Right, cur)
	}
	return cur
}
