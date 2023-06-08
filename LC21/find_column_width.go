package LC21

func findColumnWidth(grid [][]int) []int {
	res := make([]int, 100)
	maxColumn := 0
	for _, line := range grid {
		for i, num := range line {
			curLen := lengthOfNumber(num)
			if res[i] < curLen {
				res[i] = curLen
			}
			if i > maxColumn {
				maxColumn = i
			}
		}
	}

	return res[0 : maxColumn+1]

}

func lengthOfNumber(num int) int {
	if num < 0 {
		return 1 + lengthOfNumber(0-num)
	}
	if num < 10 {
		return 1
	} else {
		return 1 + lengthOfNumber(num/10)
	}

}
