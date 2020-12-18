package linkedlist

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var leftTail, rightHead, midHead, midTail, pre, cur *ListNode
	if m == 1 {
		leftTail = nil
	}
	pre = nil
	cur = head
	for i := 1; i <= n; i++ {
		if i == m-1 {
			leftTail = cur
		} else if i == m {
			midTail = cur
		}
		next := cur.Next
		if i > m {
			cur.Next = pre

		}
		pre = cur
		cur = next
	}
	midHead = pre
	rightHead = cur
	midTail.Next = rightHead
	if leftTail == nil {
		return midHead
	} else {
		leftTail.Next = midHead
		return head
	}
}
