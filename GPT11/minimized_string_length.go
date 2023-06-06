package GPT11

func MinimizedStringLength(s string) int {
	arr := make([]int, 26)
	res := 0
	for _, c := range s {
		index := c - 'a'
		if arr[index] == 0 {
			res++
			arr[index] = 1
		}
	}
	return res
}
