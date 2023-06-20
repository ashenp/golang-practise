package LC2

func removeElement(nums []int, val int) int {
	n := len(nums)
	k := 0
	for i, j := 0, n-1; i < j; {
		for nums[j] == val && j > i {
			j--
			k++
		}

		for nums[i] != val && i < j {
			i++
		}

		if i < j {
			k++
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			i++
			j--
		}
	}
	return n - k
}
