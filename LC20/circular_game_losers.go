package LC20

func circularGameLosers(n int, k int) []int {
	seats := make([]int, n)
	stop := false
	next := 0
	seats[next] = 1
	for i := 1; i <= n && stop != true; i++ {
		next = (next + i*k) % n
		if seats[next] == 1 {
			stop = true
		} else {
			seats[next] = 1
		}

	}

	res := []int{}
	for i, v := range seats {
		if v == 0 {
			res = append(res, i+1)
		}
	}
	return res
}
