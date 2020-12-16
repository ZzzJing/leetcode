package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head != nil {
		slow := head
		fast := head

		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
			if fast == slow {
				return true
			}
		}
	}

	return false
}
