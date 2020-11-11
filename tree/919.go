package tree

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	ls []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	res := CBTInserter{
		ls: []*TreeNode{nil},
	}
	currList := []*TreeNode{root}
	for len(currList) > 0 {
		nextList := []*TreeNode{}
		for i := 0; i < len(currList); i++ {
			res.ls = append(res.ls, currList[i])
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

func (this *CBTInserter) Insert(v int) int {
	new := &TreeNode{Val: v}
	this.ls = append(this.ls, new)
	tag := len(this.ls) - 1
	father := this.ls[tag/2]
	fmt.Println()
	if tag%2 == 0 {
		father.Left = new
	} else {
		father.Right = new
	}
	return father.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	if len(this.ls) == 1 {
		return nil
	}
	return this.ls[1]
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
