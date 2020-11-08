package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	list := &[]int{}
	genNodeList(root, list)
	listV := *list
	listVLen := len(listV)
	if listVLen < 2 {
		return listV[0]
	}
	min := listV[1] - listV[0]
	for i := 1; i < listVLen-1; i++ {
		curr := listV[i+1] - listV[i]
		if min > curr {
			min = curr
		}
	}
	return min
}

func genNodeList(root *TreeNode, list *[]int) {
	if root.Left != nil {
		genNodeList(root.Left, list)
	}

	*list = append(*list, root.Val)

	if root.Right != nil {
		genNodeList(root.Right, list)
	}
}
