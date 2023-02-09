package metrics

import (	
	"context"
	"os"
	"testing"
	"net/http"
	"io/ioutil"
	"encoding/json"


	"golang.org/x/oauth2"
)

var endpoint = "https://api.github.com/repos/cloudinary/cloudinary_npm"

func TestBusFactor(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	resp, error := httpClient.Get(endpoint)

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

		var bus = BusFactor(jsonRes)
		var cor = Correctness(jsonRes)
		var ramp = RampUp(jsonRes)
		var resp = ResponsiveMaintainer(jsonRes)
		var net = NetScore(cor, bus, ramp, resp, false)


		if net < 0 || net > 1 {
			t.Fatalf("Bus is out of range")
		}
	}

	defer resp.Body.Close()
}