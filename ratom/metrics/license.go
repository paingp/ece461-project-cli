package metrics

import (
	"bufio"
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"fmt"
)

var readme string

// https://stackoverflow.com/questions/71153302/how-to-set-depth-for-recursive-iteration-of-directories-in-filepath-walk-func
func walk(path string, d fs.DirEntry, err error) error {
	maxDepth := 1
	if (err != nil) {
		return err
	} 
	if d.IsDir() && strings.Count(path, string(os.PathSeparator)) > maxDepth {
		return fs.SkipDir
	} else {
		//fmt.Println(path)
		matched, _ := regexp.MatchString(`(?i)readme`, path)
		if (matched) {
			//fmt.Println("Matched: %s", path)
			check, _ := regexp.MatchString("(?i)guid", path)
			if !check {
				//fmt.Println("Matched: %s", path)
				readme = path
				//files = append(files, path)
			}
		}
	}
	return nil
}

func License(directory string) bool {
	//fmt.Println(directory)
	
	err := filepath.WalkDir(directory, walk)
	if (err != nil) {
		panic(err)
	}

	if readme != "" {
		fmt.Println(readme)
		file, err := os.Open(readme)

		if err != nil {
			panic(err)
		}
	
		defer file.Close()
	
		scanner := bufio.NewScanner(file)
		toFind := []byte(`license`)
	
		for scanner.Scan() {
			if bytes.Contains(scanner.Bytes(), toFind) {
				fmt.Println(scanner.Text())
			}
		}
	}

	return true
}
