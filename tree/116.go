package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode(left, right *Node) {
	if left == nil || right == nil {
		return
	}

	left.Next = right
	connectTwoNode(left.Left, left.Right)
	connectTwoNode(left.Right, right.Left)
	connectTwoNode(right.Left, right.Right)
}
