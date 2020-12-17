package linkedlist

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	now := head
	next := head.Next
	for next.Next != nil {
		tmp := now
		now = next
		next = next.Next
		now.Next = tmp
	}
	next.Next = now
	head.Next = nil
	return next
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	rs := reverse(head)
	head.Next = nil
	return rs
}

func reverse(node *ListNode) *ListNode {
	if node.Next == nil {
		return node
	}
	last := reverse(node.Next)
	node.Next.Next = node
	return last
}
