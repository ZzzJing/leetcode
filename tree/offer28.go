package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	currList := []*TreeNode{root.Left, root.Right}

	for len(currList) > 0 {
		length := len(currList)
		half := len(currList) / 2
		nextList := []*TreeNode{}
		for i := 0; i < length; i++ {
			if i < half {
				if currList[i] == nil || currList[length-i-1] == nil {
					if currList[i] != currList[length-i-1] {
						return false
					}
				} else {
					if currList[i].Val != currList[length-i-1].Val {
						return false
					}
				}

			}
			if currList[i] != nil {
				nextList = append(nextList, currList[i].Left)
				nextList = append(nextList, currList[i].Right)
			}
		}
		currList = nextList
	}

	return true
}
