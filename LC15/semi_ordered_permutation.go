package GPT14

func SemiOrderedPermutation(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	res := 0
	for ; left < n && nums[left] != 1; left++ {
		res++
	}

	for ; right >= 0 && nums[right] != n; right-- {
		res++
	}

	if left < n && right >= 0 && left > right {
		res--
	}
	return res
}
