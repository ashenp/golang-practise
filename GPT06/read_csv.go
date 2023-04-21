package GPT06

import (
	"bufio"
	"os"
	"strings"
)

func ReadCsv(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var res [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		token := strings.Split(line, ",")
		res = append(res, token)
	}
	return res, nil
}
