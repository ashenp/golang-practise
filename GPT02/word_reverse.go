package GPT02

func WordReverse(target string) string {
	length := len(target)
	reversed := []byte(target)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return string(reversed)
}
