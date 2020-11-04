/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package tree

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	current := []*TreeNode{root}
	for len(current) > 0 {
		next := []*TreeNode{}
		value := []int{}
		for i := 0; i < len(current); i++ {
			if current[i].Left != nil {
				next = append(next, current[i].Left)
			}
			if current[i].Right != nil {
				next = append(next, current[i].Right)
			}
			value = append(value, current[i].Val)
		}
		res = append(res, value)
		current = next
	}

	i, j := 0, len(res)-1
	for i < j {
		tmp := res[j]
		res[j] = res[i]
		res[i] = tmp
		i++
		j--
	}

	return res
}
