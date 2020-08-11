/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package tree

var count int
var node *TreeNode

func kthLargest(root *TreeNode, k int) int {
	count = 0
	Traverse(root, k)
	return node.Val
}

func Traverse(root *TreeNode, k int) {

	if root.Right != nil {
		Traverse(root.Right, k)
	}

	count++
	if count == k {
		node = root
	}

	if root.Left != nil {
		Traverse(root.Left, k)
	}
}
