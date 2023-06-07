package LC19

import "strconv"

func countSeniors(details []string) int {
	res := 0
	for _, s := range details {
		age, _ := strconv.Atoi(s[11:13])
		if age > 60 {
			res++
		}
	}
	return res
}
