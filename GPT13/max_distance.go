package GPT13

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func maxDistance(root *TreeNode, maxRes *int) int {
	if root == nil {
		return 0
	}

	if root.left == nil && root.right == nil {
		return 0
	}

	dis := max(
		height(root.left)+height(root.right),
		maxDistance(root.left, maxRes),
		maxDistance(root.right, maxRes),
	)
	if *maxRes < dis {
		*maxRes = dis
	}

	return dis
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := height(root.left)
	right := height(root.right)
	return 1 + max(left, right)
}

func max(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}

func Test() {
	root := &TreeNode{val: 5}
	root.left = &TreeNode{val: 3}
	root.right = &TreeNode{val: 7}
	root.left.left = &TreeNode{val: 2}
	root.left.right = &TreeNode{val: 4}
	root.left.left.left = &TreeNode{val: 1}
	root.right.left = &TreeNode{val: 6}
	root.right.right = &TreeNode{val: 8}
	root.right.left.right = &TreeNode{val: 9}
	root.right.left.right.left = &TreeNode{val: 10}
	root.right.left.right.right = &TreeNode{val: 11}
	var res *int = new(int)
	*res = 0
	fmt.Println("res:", maxDistance(root, res))
}
