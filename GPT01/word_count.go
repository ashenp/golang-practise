package GPT01

import (
	"bufio"
	"os"
	"strings"
)

func WordCount(filepath string) (map[string]int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	res := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasSuffix(line, ".") {
			line = line[:len(line)-1]
		}
		parts := strings.Split(line, " ")

		for _, v := range parts {
			if _, ok := res[v]; ok {
				res[v] += 1
			} else {
				res[v] = 1
			}
		}
	}
	return res, nil
}
