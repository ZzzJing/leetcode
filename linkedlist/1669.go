package linkedlist

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	node1 := list1
	var l1Pre, l1After *ListNode
	for i := 0; i < b; {
		if i == a-1 {
			l1Pre = node1
		}

		node1 = node1.Next
		i++
	}
	l1After = node1.Next

	node2 := list2
	var l2Tail *ListNode
	for node2.Next != nil {
		node2 = node2.Next
	}
	l2Tail = node2
	l1Pre.Next = list2
	l2Tail.Next = l1After
	return list1
}
