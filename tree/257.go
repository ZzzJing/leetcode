package tree

import (
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res []string

func binaryTreePaths(root *TreeNode) []string {
	res = []string{}
	if root == nil {
		return res
	}

	traverser(root, "")
	return res
}

func traverser(root *TreeNode, path string) {
	path += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		res = append(res, path)
		return
	}

	path += "->"
	if root.Left != nil {
		traverser(root.Left, path)

	}
	if root.Right != nil {
		traverser(root.Right, path)
	}
}
