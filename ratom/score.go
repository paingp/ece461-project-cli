package ratom

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/paingp/ece461-project-cli/ratom/metrics"
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

func Analyze(url string, client *http.Client) Module {
	c := metrics.Correctness()
	var urlOne string = "https://api.github.com/repos/lodash/lodash"
	resp, error := client.Get(urlOne)

	if error != nil {
		panic(error)
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(error)
		}
		bodyString := string(bodyBytes)

		resBytes := []byte(bodyString)
		var jsonRes map[string]interface{}
		_ = json.Unmarshal(resBytes, &jsonRes)

		name := jsonRes["name"].(string)
		fmt.Println(name)
	}

	m := Module{url, -1, -1, c, -1, -1, false}
	return m
}
