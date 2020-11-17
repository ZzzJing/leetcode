package stack

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	stack := []int{}
	tmp := []int{}
	res := []int{}
	node := head
	i := 0
	for node != nil {
		tmp = append(tmp, node.Val)
		for len(stack) > 0 && node.Val > tmp[stack[len(stack)-1]] {
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[pre] = node.Val
		}
		stack = append(stack, i)
		res = append(res, 0)
		i++
		node = node.Next
	}

	return res
}
