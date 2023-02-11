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
var endpoint2 = "https://api.github.com/repos/ben-ng/add"
var endpoint3 = "https://api.github.com/repos/axios/axios"
var endpoint4 = "https://api.github.com/repos/campb474/ECE368"

// Tests 1 - 5
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

		if bus < 0 || bus > 1 {
			t.Fatalf("Bus is out of range")
		}
	}

	defer resp.Body.Close()
}

func TestCorrectness(t *testing.T) {
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

		var cor = Correctness(jsonRes)


		if cor < 0 || cor > 1 {
			t.Fatalf("Correctness is out of range")
		}
	}

	

	defer resp.Body.Close()
}

func TestRampUp(t *testing.T) {
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

		var ramp = RampUp(jsonRes, 12)


		if ramp < 0 || ramp > 1 {
			t.Fatalf("Ramp Up is out of range")
		}
	}

	defer resp.Body.Close()
}

func TestResponsiveMaintainer(t *testing.T) {
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

		var resp = ResponsiveMaintainer(jsonRes)

		if resp < 0 || resp > 1 {
			t.Fatalf("Responsive Maintainer is out of range")
		}
	}

	

	defer resp.Body.Close()
}


func TestNetScore(t *testing.T) {
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
		var ramp = RampUp(jsonRes, 12)
		var resp = ResponsiveMaintainer(jsonRes)
		var net = NetScore(cor, bus, ramp, resp, false)


		if net < 0 || net > 1 {
			t.Fatalf("Net Score is out of range")
		}
	}

	

	defer resp.Body.Close()
}

// Tests 6 - 10
func TestBusFactor2(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	resp, error := httpClient.Get(endpoint2)

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

		if bus < 0 || bus > 1 {
			t.Fatalf("Bus is out of range")
		}
	}

	defer resp.Body.Close()
}

func TestCorrectness2(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	resp, error := httpClient.Get(endpoint2)

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

		var cor = Correctness(jsonRes)


		if cor < 0 || cor > 1 {
			t.Fatalf("Correctness is out of range")
		}
	}

	

	defer resp.Body.Close()
}

func TestRampUp2(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	resp, error := httpClient.Get(endpoint2)

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

		var ramp = RampUp(jsonRes, 12)


		if ramp < 0 || ramp > 1 {
			t.Fatalf("Ramp Up is out of range")
		}
	}

	defer resp.Body.Close()
}

func TestResponsiveMaintainer2(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	resp, error := httpClient.Get(endpoint2)

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

		var resp = ResponsiveMaintainer(jsonRes)

		if resp < 0 || resp > 1 {
			t.Fatalf("Responsive Maintainer is out of range")
		}
	}

	

	defer resp.Body.Close()
}


func TestNetScore2(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	resp, error := httpClient.Get(endpoint2)

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
		var ramp = RampUp(jsonRes, 12)
		var resp = ResponsiveMaintainer(jsonRes)
		var net = NetScore(cor, bus, ramp, resp, false)


		if net < 0 || net > 1 {
			t.Fatalf("Net Score is out of range")
		}
	}

	

	defer resp.Body.Close()
}

// Tests 11 - 15
func TestBusFactor3(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	resp, error := httpClient.Get(endpoint3)

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

		if bus < 0 || bus > 1 {
			t.Fatalf("Bus is out of range")
		}
	}

	defer resp.Body.Close()
}

func TestCorrectness3(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	resp, error := httpClient.Get(endpoint3)

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

		var cor = Correctness(jsonRes)


		if cor < 0 || cor > 1 {
			t.Fatalf("Correctness is out of range")
		}
	}

	

	defer resp.Body.Close()
}

func TestRampUp3(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	resp, error := httpClient.Get(endpoint3)

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

		var ramp = RampUp(jsonRes, 12)


		if ramp < 0 || ramp > 1 {
			t.Fatalf("Ramp Up is out of range")
		}
	}

	defer resp.Body.Close()
}

func TestResponsiveMaintainer3(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	resp, error := httpClient.Get(endpoint3)

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

		var resp = ResponsiveMaintainer(jsonRes)

		if resp < 0 || resp > 1 {
			t.Fatalf("Responsive Maintainer is out of range")
		}
	}

	

	defer resp.Body.Close()
}


func TestNetScore3(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	resp, error := httpClient.Get(endpoint3)

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
		var ramp = RampUp(jsonRes, 12)
		var resp = ResponsiveMaintainer(jsonRes)
		var net = NetScore(cor, bus, ramp, resp, false)


		if net < 0 || net > 1 {
			t.Fatalf("Net Score is out of range")
		}
	}

	defer resp.Body.Close()
}

// Test 16

func TestBusFactor_private(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	resp, error := httpClient.Get(endpoint4)

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

		if bus < 0 || bus > 1 {
			t.Fatalf("Bus is out of range")
		}
	}

	defer resp.Body.Close()
}