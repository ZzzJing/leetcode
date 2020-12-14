package tree

func recoverTree(root *TreeNode) {
	var x, y, pre *TreeNode
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node.Left != nil {
			traverse(node.Left)
		}
		if pre == nil {
			pre = node
		} else {
			if pre.Val > node.Val {

				y = node
				if x == nil {
					x = pre
				} else {
					return
				}
			}
			pre = node
		}

		if node.Right != nil {
			traverse(node.Right)
		}
	}
	traverse(root)
	if x != nil && y != nil {
		x.Val, y.Val = y.Val, x.Val
	}

}
