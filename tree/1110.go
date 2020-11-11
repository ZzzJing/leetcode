package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	res := []*TreeNode{}
	if root == nil {
		return res
	}
	deleteMap := make(map[int]bool)
	for _, v := range to_delete {
		deleteMap[v] = true
	}

	var afterTrave func(node, father *TreeNode)
	afterTrave = func(node, father *TreeNode) {
		if node.Left != nil {
			afterTrave(node.Left, node)
		}

		if node.Right != nil {
			afterTrave(node.Right, node)
		}

		if _, exists := deleteMap[node.Val]; exists {
			if node.Left != nil {
				res = append(res, node.Left)
			}
			if node.Right != nil {
				res = append(res, node.Right)
			}
			if father != nil && father.Left == node {
				father.Left = nil
			}
			if father != nil && father.Right == node {
				father.Right = nil
			}
		}
	}
	if _, exists := deleteMap[root.Val]; !exists {
		res = append(res, root)
	}
	afterTrave(root, nil)
	return res
}
