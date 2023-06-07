package LC16

func BuyChoco(prices []int, money int) int {
	min1, min2 := findMinTwo(prices)

	if min1+min2 > money {
		return money
	}
	return money - (min1 + min2)

}

func findMinTwo(prices []int) (int, int) {
	min1, min2 := prices[0], prices[1]
	if min1 > min2 {
		min1, min2 = min2, min1
	}

	for i := 2; i < len(prices); i++ {
		if prices[i] < min1 {
			min2 = min1
			min1 = prices[i]
		} else if prices[i] < min2 {
			min2 = prices[i]
		}
	}

	return min1, min2
}
