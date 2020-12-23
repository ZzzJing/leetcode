package tree

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	var res *TreeNode
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node.Left != nil {
			helper(node.Left)
		}

		count++
		if count == k {
			res = node
			return
		}

		if node.Right != nil {
			helper(node.Right)
		}
	}

	helper(root)
	if res != nil {
		return res.Val
	} else {
		return 0
	}
}
