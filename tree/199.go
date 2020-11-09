package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	currList := []*TreeNode{root}
	for len(currList) > 0 {
		nextList := []*TreeNode{}
		length := len(currList)
		res = append(res, currList[length-1].Val)
		for i := 0; i < length; i++ {
			if currList[i].Left != nil {
				nextList = append(nextList, currList[i].Left)
			}
			if currList[i].Right != nil {
				nextList = append(nextList, currList[i].Right)
			}
		}
		currList = nextList
	}

	return res
}
