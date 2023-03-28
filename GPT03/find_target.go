package GPT02

func FindTarget(nums []int, target int) (int, int) {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		_, ok := dict[diff]
		if ok {
			return dict[diff], i
		} else {
			dict[nums[i]] = i
		}
	}
	return -1, -1
}
