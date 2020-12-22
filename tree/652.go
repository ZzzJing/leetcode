package tree

import (
	"fmt"
	"strconv"
)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	treeMap := map[string]int{}

	var helper func(node *TreeNode) string
	helper = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		left:= helper(node.Left)
		right := helper(node.Right)
		nodeString := left + "," + right + "," + strconv.Itoa(node.Val)
		if v, exist := treeMap[nodeString]; exist {
			fmt.Println(nodeString, v)
			if v == 1 {
				res = append(res, node)
			}
			treeMap[nodeString]++
		} else {
			treeMap[nodeString] = 1
		}

		return nodeString
	}

	helper(root)
	return res
}
