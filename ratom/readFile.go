package ratom

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filePath string) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var urls []string

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		urls = append(urls, scanner.Text())
	}
}
