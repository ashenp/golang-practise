package GPT06

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

func WordSorting(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	matrix := make([][]string, 26)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		offset := line[0] - 'a'
		matrix[offset] = append(matrix[offset], strings.TrimRight(line, "\n"))
	}

	var buffer bytes.Buffer
	for _, v := range matrix {
		for _, v2 := range v {
			buffer.WriteString(v2)
			buffer.WriteString("\n")
		}
	}

	return buffer.String(), nil
}
