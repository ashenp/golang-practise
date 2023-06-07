package LC17

func MinLength(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	b := []byte{}

	for _, c := range s {
		if c == 'B' && len(b) > 0 && b[len(b)-1] == 'A' {
			b = b[:len(b)-1]
		} else if c == 'D' && len(b) > 0 && b[len(b)-1] == 'C' {
			b = b[:len(b)-1]
		} else {
			b = append(b, byte(c))
		}
	}
	return len(b)
}
