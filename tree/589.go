/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
package tree

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	list := []int{}
	list = append(list, root.Val)

	if len(root.Children) > 0 {
		for _, v := range root.Children {
			list = append(list, preorder(v)...)
		}
	}
	return list
}
