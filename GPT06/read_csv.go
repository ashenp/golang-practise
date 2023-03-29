package GPT06

import (
	"bufio"
	"os"
	"strings"
)

func ReadCsv(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		return nil
	}

	var res [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		token := strings.Split(line, ",")
		res = append(res, token)
	}
	return res
}
