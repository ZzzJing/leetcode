package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	l1 := findLeafNode(root1)
	l2 := findLeafNode(root2)

	if len(l1) != len(l2) {
		return false
	}
	for i, _ := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true

}

func findLeafNode(root *TreeNode) []int {
	leafList := []int{}
	if root == nil {
		return leafList
	}
	if root.Left == nil && root.Right == nil {
		leafList = append(leafList, root.Val)
		return leafList
	}

	leafList = append(leafList, findLeafNode(root.Left)...)

	leafList = append(leafList, findLeafNode(root.Right)...)

	return leafList
}
