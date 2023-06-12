package LC23

func evenOddBit(n int) []int {

	index := 0
	res := []int{0, 0}
	for ; n > 0; n = n / 2 {
		c := n % 2
		if c == 1 {
			res[index] = res[index] + 1
		}
		index = 1 - index
	}

	return res
}
