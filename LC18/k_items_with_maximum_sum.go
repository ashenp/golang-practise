package LC18

func KItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if k < numOnes {
		return k
	}

	res := 0
	res += numOnes
	k -= numOnes

	if k <= numZeros {
		return res
	}
	k -= numZeros

	if k >= 0 {
		res -= k
	}
	return res
}
