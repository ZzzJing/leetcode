package tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var helper func(left, right int) *TreeNode
	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		max := left
		for i := left + 1; i <= right; i++ {
			if nums[i] > nums[max] {
				max = i
			}
		}

		return &TreeNode{
			Val:   nums[max],
			Left:  helper(left, max-1),
			Right: helper(max+1, right),
		}
	}
	return helper(0, len(nums)-1)
}
