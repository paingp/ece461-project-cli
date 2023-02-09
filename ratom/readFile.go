package ratom

import (
	"bufio"
	"os"
)

func ReadFile(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var urls []string

	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	return urls
}
