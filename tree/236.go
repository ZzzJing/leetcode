package tree

func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := lowestCommonAncestor236(root.Left, p, q)
	right := lowestCommonAncestor236(root.Right, p, q)

	if left == nil && right == nil {
		return nil
	}

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	return root
}
