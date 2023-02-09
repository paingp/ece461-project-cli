package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	//"sort"

	"github.com/paingp/ece461-project-cli/ratom"

	//"time"
	"github.com/joho/godotenv"
	// "github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func main() {
	// Read input file
	filePath := os.Args[1]

	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	//var urls []string

	err = godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	// Create a token to and HTTP client to access the GitHub API
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	err = os.Mkdir("temp", 0770)		// read, write, and execute access for file owner and group owner
	if err != nil {
		panic("Error creating temp directory")
	}

	var modules []ratom.Module

	// Read file line by line
	for scanner.Scan() {
		url := scanner.Text()
		module := ratom.Analyze(url, httpClient)
		fmt.Println(module)
		modules = append(modules, module)
		//urls = append(urls, url)
	}

	os.RemoveAll("temp")

	// sort.SliceStable(modules, func(i, j int) bool {
	// 	return &modules[i].netScore < modules[j].netScore
	// })

	// Load .env file
	//
	// requestURL := "https://api.github.com/repos/lencx/ChatGPT"
	// resp, error := httpClient.Get(requestURL)

	// if error != nil {
	// 	panic(error)
	// }

	// defer resp.Body.Close()

	// body, error := ioutil.ReadAll(resp.Body)
	// if error != nil {
	// 	panic(error)
	// }

	// fmt.Println(string(body))

	// // Use GitHub GrapQL API via githubv4 library

	// client := githubv4.NewClient(httpClient)
	// fmt.Println(os.Getenv("GITHUB_TOKEN"))

	// // Variable to store results of GraphQL query
	// var q struct {
	// 	Viewer struct {
	// 		Login     string
	// 		CreatedAt time.Time
	// 	}
	// }

	// error = client.Query(context.Background(), &q, nil)
	// if error != nil {
	// 	panic(error)
	// }

	// fmt.Println(q.Viewer.Login)
	// fmt.Println(q.Viewer.CreatedAt)
}
