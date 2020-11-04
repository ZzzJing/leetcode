/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package tree

func averageOfLevels(root *TreeNode) []float64 {
	nextLevel := []*TreeNode{}
	if root != nil {
		nextLevel = append(nextLevel, root)
	}
	res := []float64{}
	for len(nextLevel) > 0 {
		currentLevel := nextLevel
		nextLevel = []*TreeNode{}
		sum := 0
		for i := 0; i < len(currentLevel); i++ {
			sum += currentLevel[i].Val
			if currentLevel[i].Left != nil {
				nextLevel = append(nextLevel, currentLevel[i].Left)
			}
			if currentLevel[i].Right != nil {
				nextLevel = append(nextLevel, currentLevel[i].Right)
			}
		}

		res = append(res, float64(sum)/float64(len(currentLevel)))
	}

	return res
}
