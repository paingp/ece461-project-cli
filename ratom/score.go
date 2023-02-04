package ratom

import (
	"github.com/paingp/ece461-project-cli/ratom/metrics"
	"net/http"
	"fmt"
	"encoding/json"
	//"io/ioutil"
)

type Module struct {
	url         string
	netScore    float32
	rampUp      float32
	correctness float32
	busFactor   float32
	respMaint   float32
	license     bool
}

func formatUrl(url string) string {
	return "";
}

func Analyze(url string, client *http.Client) Module {
	resp, error := client.Get("https://api.github.com/repos/lencx/ChatGPT")

	if error != nil {
		panic(error)
	}

	defer resp.Body.Close()

	var j interface{}
	error = json.NewDecoder(resp.Body).Decode(&j)
	if error != nil {
			panic(error)
	}
	fmt.Printf("%s", j)
	fmt.Println(resp)

	// defer resp.Body.Close()

	// body, error := ioutil.ReadAll(resp.Body)
	// if error != nil {
	// 	panic(error)
	// }

	// fmt.Println(string(body))

	if error != nil {
		panic(error)
	}

	c := metrics.Correctness()
	m := Module{url, -1, -1, c, -1, -1, false}
	return m
}
