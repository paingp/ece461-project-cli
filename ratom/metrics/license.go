package metrics

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
	
	var text []string		// store array of strings of lines that contain the word license

	err := filepath.WalkDir(directory, walk)
	if err != nil {
		Functions = append(Functions, "Error while trying to find ReadMe in " + directory)
		return false
	}

	if readme == "" {
		Functions = append(Functions, "ReadMe could be found in " + directory)
		return false
	}

	file, err := os.Open(readme)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`(?i)license`)

	for scanner.Scan() {
		line := scanner.Bytes()

		if (re.Match(line)) {
			//fmt.Println(string(line))
			text = append(text, string(line))
		}
	}

	licenses := [9]string{"MIT", "LGPLv2.1", "Expat", "X11", "MPL-2.0", "Mozilla Public", "Artistic License 2", "GPLv2", "GPLv3"}

	for i := 0; i < len(text); i++ {
		//fmt.Println(text[i])
		for j := 0; j < len(licenses); j++ {
			re = regexp.MustCompile("(?i)" + licenses[j])
			if (re.MatchString(text[i])) {
				return true
			}
		}
	}

	return false
}
