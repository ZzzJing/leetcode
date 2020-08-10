/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
package tree

func postorder(root *Node) []int {
	list := []int{}
	if root == nil {
		return list
	}

	if len(root.Children) != 0 {
		for _, v := range root.Children {
			list = append(list, postorder(v)...)
		}
	}

	list = append(list, root.Val)
	return list
}
