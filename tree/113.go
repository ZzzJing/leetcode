package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	var getSum func(root *TreeNode, sumList []int, sumCount int)
	getSum = func(root *TreeNode, sumList []int, sumCount int) {
		sumCount = sumCount + root.Val
		sumList = append(sumList, root.Val)
		if sumCount == sum {
			if root.Left == nil && root.Right == nil {
				tmp := make([]int, len(sumList))
				copy(tmp, sumList)
				res = append(res, tmp)
			}

		}
		if root.Left != nil {
			getSum(root.Left, sumList, sumCount)
		}
		if root.Right != nil {
			getSum(root.Right, sumList, sumCount)
		}
		sumList = sumList[:len(sumList)-1]
	}
	getSum(root, []int{}, 0)
	return res
}
