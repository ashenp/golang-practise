package LC25

func Test() {
	wordPattern("abba", "dog dog dog dog")
}

func wordPattern(pattern string, s string) bool {
	m := make(map[byte]string)
	t := make(map[string]bool)

	left := 0
	index := -1
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			w := s[left:i]
			left = i + 1

			index++
			if index >= len(pattern) {
				return false
			}
			c := pattern[index]
			if v, ok := m[c]; ok {
				if v != w {
					return false
				}
			} else {
				if _, ok := t[w]; ok {
					return false
				} else {
					m[c] = w
					t[w] = true
				}
			}
		}
	}
	return true
}
