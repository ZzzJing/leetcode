package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var father *ListNode
	node := head
	for node.Next != nil && n > 1 {
		node = node.Next
		n--
	}

	if n == 1 {
		new := head
		for node.Next != nil {
			father = new
			new = new.Next
			node = node.Next
		}
		if father == nil {
			return head.Next
		}
		father.Next = new.Next
	}

	return head
}
