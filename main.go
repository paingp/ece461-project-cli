package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"

	//"sort"

	"github.com/paingp/ece461-project-cli/ratom"
	"github.com/paingp/ece461-project-cli/ratom/metrics"

	//"time"
	"github.com/joho/godotenv"
	// "github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// Function to get the current line number - citation needed
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
	error := godotenv.Load(".env")
	if error != nil {
		panic("Error loading .env file")
	}

	// Create a token to and HTTP client to access the GitHub API
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))

	if err != nil {
		fmt.Println("Error during string to int conversion")
		return
	}

	var modules []ratom.Module

	for scanner.Scan() {
		url := scanner.Text()
		metrics.Functions = append(metrics.Functions, "For URL: "+url)
		module := ratom.Analyze(url, httpClient)
		lineNumb := metrics.File_line()
		metrics.Functions = append(metrics.Functions, "Function: ratom.Analyze called on main.go at line "+lineNumb)
		//fmt.Println(module)
		modules = append(modules, module)
		//urls = append(urls, url)
	}

	// Sorting the modules by decreasing net score
	sort.SliceStable(modules, func(i, j int) bool {
		return modules[i].NetScore > modules[j].NetScore
	})
	lineNumb := metrics.File_line()
	metrics.Functions = append(metrics.Functions, "Function: SliceStable called on main.go at line "+lineNumb)

	for a := 0; a < len(modules); a++ {
		var b = 0
		if modules[a].License {
			b = 1
		}
		fmt.Printf("{\"URL\":%s, \"NET_SCORE\":%1.1f, \"RAMP_UP_SCORE\":%1.1f, \"CORRECTIVENESS_SCORE\":%1.1f, \"BUS_FACTOR_SCORE\":%1.1f, \"RESPONSIVENESS_MAINTAINER_SCORE\":%1.1f, \"LICENSE SCORE\":%d}\n", modules[a].Url, modules[a].NetScore, modules[a].RampUp, modules[a].Correctness, modules[a].BusFactor, modules[a].RespMaint, b)
		//fmt.Println(modules[a])
	}

	if logLevel == 1 {
		ratom.LoggerVerbOne(modules)
	} else if logLevel == 2 {
		ratom.LoggerVerbTwo(modules)
	}

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
