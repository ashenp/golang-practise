package LC22

func diagonalPrime(nums [][]int) int {
	res := 0
	n := len(nums)
	for i := 0; i < n; i = i + 1 {

		j := n - 1 - i
		x := nums[i][i]
		y := nums[i][j]

		if isPrime(x) {
			if res < x {
				res = x
			}
		}

		if isPrime(y) {
			if res < y {
				res = y
			}
		}

	}

	return res
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true

}
