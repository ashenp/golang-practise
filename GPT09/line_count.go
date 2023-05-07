package GPT09

import (
	"bufio"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sync"
)

type FileInfo struct {
	filePath string
	lines    int
}

func scanFiles(dirPath string, filesChan chan string) {
	filepath.Walk(dirPath, func(filePath string, info fs.FileInfo, err error) error {
		if !info.IsDir() && path.Ext(filePath) == ".go" {
			filesChan <- filePath
		}
		return nil
	})
	close(filesChan)
}

func countLines(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		return 0
	}
	return lineCount
}

func countDirLines(filesChan chan string, linesChan chan FileInfo, wg *sync.WaitGroup) {
	defer wg.Done()

	for filePath := range filesChan {
		count := countLines(filePath)
		linesChan <- FileInfo{
			filePath: filePath,
			lines:    count,
		}
	}
}

func CountLinesConcurrently(dirPath string, numCores int) int {
	filesChan := make(chan string, 100)

	go scanFiles(dirPath, filesChan)

	linesChan := make(chan FileInfo, 100)

	var wg sync.WaitGroup

	for i := 0; i < numCores; i++ {
		wg.Add(1)
		go countDirLines(filesChan, linesChan, &wg)
	}

	go func() {
		wg.Wait()
		close(linesChan)
	}()

	totalLines := 0
	for fileInfo := range linesChan {
		totalLines += fileInfo.lines
	}
	return totalLines
}
