package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		if p == q {
			return true
		} else {
			return false
		}
	}
	if p.Val != q.Val {
		return false
	}
	currListP := []*TreeNode{p}
	currListQ := []*TreeNode{q}
	for len(currListP) > 0 || len(currListQ) > 0 {
		if len(currListP) != len(currListQ) {
			return false
		}
		nextListP := []*TreeNode{}
		nextListQ := []*TreeNode{}
		for i := 0; i < len(currListP); i++ {
			if currListP[i] == nil || currListQ[i] == nil {
				if currListP[i] != currListQ[i] {
					return false
				}
			} else {
				if currListP[i].Val != currListQ[i].Val {
					return false
				}
			}

			if currListP[i] != nil {
				nextListP = append(nextListP, currListP[i].Left)
				nextListP = append(nextListP, currListP[i].Right)
			}

			if currListP[i] != nil {
				nextListQ = append(nextListQ, currListQ[i].Left)
				nextListQ = append(nextListQ, currListQ[i].Right)
			}

		}

		currListP = nextListP
		currListQ = nextListQ
	}

	return true
}
