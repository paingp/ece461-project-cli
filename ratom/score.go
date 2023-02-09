package ratom

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"os/exec"
	"strings"

	"github.com/paingp/ece461-project-cli/ratom/metrics"
	"github.com/go-git/go-git/v5"
)

var GITHUB_TOKEN string

type Module struct {
	url         string
	netScore    float32
	rampUp      float32
	correctness float32
	busFactor   float32
	respMaint   float32
	license     bool
}

func getEndpoint(url string) string {
	endpoint := "Invalid"
	//if (strings.Contains(url, "github"))
	before, after, found := strings.Cut(url, "www")
	if found {
		endpoint = before + "registry" + after
		endpoint = strings.Replace(endpoint, "com", "org", 1)
		endpoint = strings.Replace(endpoint, "package/", "", 1)

		resp, err := http.Get(endpoint)

		if err != nil {
			return "Invalid"
		}

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				panic(err)
			}
			bodyString := string(bodyBytes)

			resBytes := []byte(bodyString)
			var npmRes map[string]interface{}
			_ = json.Unmarshal(resBytes, &npmRes)

			bugs := npmRes["bugs"].(map[string]interface{})
			endpoint = bugs["url"].(string)
			Clone(strings.Replace(endpoint, "/issues", ".git", 1))
			endpoint = endpoint[:8] + "api." + strings.Replace(endpoint[8:], "/", "/repos/", 1)
			endpoint = endpoint[:len(endpoint)-7]
		}

	} else {
		Clone(url + ".git")
		index := strings.Index(url, "github")
		if index != -1 {
			//fmt.Print(url[index:])
			endpoint = "https://api." + strings.Replace(url[index:], "/", "/repos/", 1)
		}
		//endpoint = subStrings[0] + "api." + subStrings[1]
	}
	return endpoint
}

func GetToken() string {
	return os.Getenv("GITHUB_TOKEN")
}

func Clone(repo string) {

	// Temp directory to clone the repository
	if GITHUB_TOKEN == "" {
		GITHUB_TOKEN = GetToken()
	}

	dir, err := os.MkdirTemp("temp", "repo")
	//fmt.Println(url)

	if err != nil {
		panic(err)
	}

	defer os.RemoveAll(dir)
	log.Println(dir)
	log.Println(repo)

	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: repo,
		SingleBranch: true,
		Depth: 1,
	})
		
	if err != nil {
		log.Fatal(err)
	}

	// cmd := exec.Command("ls | grep -i readme", dir)
	// out, err := cmd.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(string(out))
}

func Analyze(url string, client *http.Client) Module {

	var correctnessScore float32
	var busFactor float32
	var rampUp float32
	var responsiveMaintainer float32
	var license bool
	var netScore float32

	oldURL := url
	url = getEndpoint((url))
	resp, error := client.Get(url)

	if error != nil {
		panic(error)
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(error)
		}
		bodyString := string(bodyBytes)

		// Citation needed for this
		resBytes := []byte(bodyString)
		var jsonRes map[string]interface{}
		_ = json.Unmarshal(resBytes, &jsonRes)

		//name := jsonRes["id"].(float64)
		correctnessScore = metrics.Correctness(jsonRes)
		busFactor = metrics.BusFactor(jsonRes)
		rampUp = metrics.RampUp(jsonRes)
		responsiveMaintainer = metrics.ResponsiveMaintainer(jsonRes)
		license = metrics.License(jsonRes)
		netScore = metrics.NetScore(correctnessScore, busFactor, rampUp, responsiveMaintainer, license)
	}

	m := Module{oldURL, netScore, rampUp, correctnessScore, busFactor, responsiveMaintainer, license}
	return m
}
