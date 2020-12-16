package linkedlist

func kthToLast(head *ListNode, k int) int {
	tmp := head
	for k > 1 {
		tmp = tmp.Next
		k--
	}

	res := head
	for tmp.Next != nil {
		res = res.Next
		tmp = tmp.Next
	}

	return res.Val
}
