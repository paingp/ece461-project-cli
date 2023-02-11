package ratom

import (
	"context"
	"encoding/json"
	//"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/shurcooL/githubv4"

	"github.com/go-git/go-git/v5"
	"github.com/paingp/ece461-project-cli/ratom/metrics"
)

var GITHUB_TOKEN string

type Module struct {
	Url         string
	NetScore    float32
	RampUp      float32
	Correctness float32
	BusFactor   float32
	RespMaint   float32
	License     bool
}

func getGithubUrl(url string) string {
	before, after, found := strings.Cut(url, "www")
	if found {
		npmEndpoint := before + "registry" + after
		npmEndpoint = strings.Replace(npmEndpoint, "com", "org", 1)
		npmEndpoint = strings.Replace(npmEndpoint, "package/", "", 1)

		resp, err := http.Get(npmEndpoint)

		if err != nil {
			return "Invalid"
		}

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)

			if err != nil {
				panic(err)
			}
			bodyString := string(bodyBytes)

			resBytes := []byte(bodyString)
			var npmRes map[string]interface{}
			_ = json.Unmarshal(resBytes, &npmRes)

			bugs := npmRes["bugs"].(map[string]interface{})
			npmEndpoint = bugs["url"].(string)
			url = strings.Replace(npmEndpoint, "/issues", "", 1)
		}
	}
	return url
}

func getEndpoint(url string) string {
	index := strings.Index(url, "github")
	url = "https://api." + strings.Replace(url[index:], "/", "/repos/", 1)
	return url
}

func GetToken() string {
	return os.Getenv("GITHUB_TOKEN")
}

func Clone(repo string) string {

	// Temp directory to clone the repository
	if GITHUB_TOKEN == "" {
		GITHUB_TOKEN = GetToken()
	}

	lastIdx := strings.LastIndex(repo, "/")
	//fmt.Println(repo[lastIdx + 1:])
	dir := "temp/" + repo[lastIdx + 1:]

	err := os.MkdirAll(dir, 0777)

	if err != nil {
		panic(err)
		return ""
	}

	//defer os.RemoveAll(dir)
	// log.Println(dir)
	// log.Println(repo)

	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: repo + ".git",
		SingleBranch: true,
		Depth: 1,
	})

	if err != nil {
		log.Fatal(err)
		return "err"
	}
	return dir
}

func Analyze(url string, client *http.Client) Module {

	var busFactor float32
	var responsiveMaintainer float32
	var correctnessScore float32
	var rampUp float32
	var license bool
	var netScore float32

	//oldURL := url

	gitUrl := getGithubUrl(url)
	//fmt.Println(gitUrl)
	dir := Clone(gitUrl)

	endpoint := getEndpoint(gitUrl)
	lineNumb := metrics.File_line()
	metrics.Functions = append(metrics.Functions, "Function: getEndpoint called on score.go at line "+lineNumb)

	resp, error := client.Get(endpoint)

	if error != nil {
		panic(error)
	}
	
	//fmt.Println(Data.Repository.CommitComments.TotalCount)

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			panic(error)
		}
		bodyString := string(bodyBytes)

		// Citation needed for this
		resBytes := []byte(bodyString)
		var jsonRes map[string]interface{}
		_ = json.Unmarshal(resBytes, &jsonRes)

		// GRAPH QL 
		owner_map := jsonRes["owner"].(map[string]interface{})

		var Data struct {
			Viewer struct {
				Login string
			}
			Repository struct {
				CommitComments struct {
					TotalCount int
				}
			} `graphql:"repository(owner: $owner, name: $name)"`
		}
	
		variables := map[string]interface{}{
			"owner": githubv4.String(owner_map["login"].(string)),
			"name":  githubv4.String(jsonRes["name"].(string)),
		}
	
		graphQLClient := githubv4.NewClient(client)
		error = graphQLClient.Query(context.Background(), &Data, variables)
		if error != nil {
			panic(error)
		}
		
		//name := jsonRes["id"].(float64)
		correctnessScore = metrics.Correctness(jsonRes)
		lineNumb := metrics.File_line()
		metrics.Functions = append(metrics.Functions, "Function: metrics.Correctness called on score.go at line "+lineNumb)

		busFactor = metrics.BusFactor(jsonRes)
		lineNumb = metrics.File_line()
		metrics.Functions = append(metrics.Functions, "Function: metrics.BusFactor called on score.go at line "+lineNumb)

		rampUp = metrics.RampUp(jsonRes, Data.Repository.CommitComments.TotalCount)
		lineNumb = metrics.File_line()
		metrics.Functions = append(metrics.Functions, "Function: metrics.RampUp called on score.go at line "+lineNumb)

		responsiveMaintainer = metrics.ResponsiveMaintainer(jsonRes)
		lineNumb = metrics.File_line()
		metrics.Functions = append(metrics.Functions, "Function: metrics.ResponsiveMaintainer called on score.go at line "+lineNumb)


		license = metrics.License(dir)
		lineNumb = metrics.File_line()
		metrics.Functions = append(metrics.Functions, "Function: metrics.License called on score.go at line "+lineNumb)

		netScore = metrics.NetScore(correctnessScore, busFactor, rampUp, responsiveMaintainer, license)
		lineNumb = metrics.File_line()
		metrics.Functions = append(metrics.Functions, "Function: metrics.NetScore called on score.go at line "+lineNumb)
	}

	defer resp.Body.Close()

	m := Module{url, netScore, rampUp, correctnessScore, busFactor, responsiveMaintainer, license}
	return m
}
