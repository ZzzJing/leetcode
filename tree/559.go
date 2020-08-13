/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
package tree

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	max := 0
	if len(root.Children) > 0 {
		for _, v := range root.Children {
			dep := maxDepth(v)
			if dep > max {
				max = dep
			}
		}
	}

	return max + 1
}
