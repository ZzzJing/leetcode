package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return false
	}
	find := false
	count := make(map[int]struct{})
	var trave func(node *TreeNode)
	trave = func(node *TreeNode) {
		if _, exist := count[k-node.Val]; exist {
			find = true
			return
		}
		count[node.Val] = struct{}{}
		if !find && node.Left != nil {
			trave(node.Left)
		}

		if !find && node.Right != nil {
			trave(node.Right)
		}
	}

	trave(root)
	return find
}
