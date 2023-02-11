package ratom

<<<<<<< HEAD
<<<<<<<< HEAD:ratom/ratom_test.go
=======
>>>>>>> 7d64c38 (Updating scoring and testing)
import (
	"testing"
	"fmt"
	"golang.org/x/oauth2"
	"os"
	"context"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/paingp/ece461-project-cli/ratom/metrics"
)

var file = "https://github.com/lodash/lodash"

func TestAnalyze(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	module := Analyze(file, httpClient)
<<<<<<< HEAD
	os.RemoveAll("temp")
=======
>>>>>>> 7d64c38 (Updating scoring and testing)
	var modules []Module
	modules = append(modules, module)
	LoggerVerbOne(modules)
	LoggerVerbTwo(modules)

	if(module.License == false) {
		fmt.Printf("No License")
	}
}

var endpoint = "https://api.github.com/repos/cloudinary/cloudinary_npm"
var endpoint2 = "https://api.github.com/repos/ben-ng/add"
var endpoint3 = "https://api.github.com/repos/axios/axios"
var endpoint4 = "https://api.github.com/repos/campb474/ECE368"

// Tests 1 - 6
func TestBusFactor(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
<<<<<<< HEAD
========
// import (	
// 	"context"
// 	"os"
// 	"testing"
// 	"net/http"
// 	"io/ioutil"
// 	"encoding/json"


// 	"golang.org/x/oauth2"
// )

// var endpoint = "https://api.github.com/repos/cloudinary/cloudinary_npm"
// var endpoint2 = "https://api.github.com/repos/ben-ng/add"
// var endpoint3 = "https://api.github.com/repos/axios/axios"
// var endpoint4 = "https://api.github.com/repos/campb474/ECE368"

// // Tests 1 - 6
// func TestBusFactor(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go

// 	resp, error := httpClient.Get(endpoint)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
		var bus = metrics.BusFactor(jsonRes)
========
// 		var bus = BusFactor(jsonRes)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go

// 		if bus < 0 || bus > 1 {
// 			t.Fatalf("Bus is out of range")
// 		}
// 	}

// 	defer resp.Body.Close()
// }

// func TestCorrectness(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	resp, error := httpClient.Get(endpoint)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
		var cor = metrics.Correctness(jsonRes)
========
// 		var cor = Correctness(jsonRes)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go


// 		if cor < 0 || cor > 1 {
// 			t.Fatalf("Correctness is out of range")
// 		}
// 	}

	

// 	defer resp.Body.Close()
// }

// func TestRampUp(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	resp, error := httpClient.Get(endpoint)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
		var ramp = metrics.RampUp(jsonRes, 3)
========
// 		var ramp = RampUp(jsonRes, 3)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go


// 		if ramp < 0 || ramp > 1 {
// 			t.Fatalf("Ramp Up is out of range")
// 		}
// 	}

// 	defer resp.Body.Close()
// }

// func TestResponsiveMaintainer(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	resp, error := httpClient.Get(endpoint)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
		var resp = metrics.ResponsiveMaintainer(jsonRes)
========
// 		var resp = ResponsiveMaintainer(jsonRes)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go

// 		if resp < 0 || resp > 1 {
// 			t.Fatalf("Responsive Maintainer is out of range")
// 		}
// 	}

	

// 	defer resp.Body.Close()
// }


// func TestNetScore(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	resp, error := httpClient.Get(endpoint)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
=======

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

		var bus = metrics.BusFactor(jsonRes)

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

		var cor = metrics.Correctness(jsonRes)


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

		var ramp = metrics.RampUp(jsonRes, 3)


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

		var resp = metrics.ResponsiveMaintainer(jsonRes)

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

>>>>>>> 7d64c38 (Updating scoring and testing)
		var bus = metrics.BusFactor(jsonRes)
		var cor = metrics.Correctness(jsonRes)
		var ramp = metrics.RampUp(jsonRes, 20)
		var resp = metrics.ResponsiveMaintainer(jsonRes)
		var net = metrics.NetScore(cor, bus, ramp, resp, false)
<<<<<<< HEAD
========
// 		var bus = BusFactor(jsonRes)
// 		var cor = Correctness(jsonRes)
// 		var ramp = RampUp(jsonRes, 20)
// 		var resp = ResponsiveMaintainer(jsonRes)
// 		var net = NetScore(cor, bus, ramp, resp, false)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go


// 		if net < 0 || net > 1 {
// 			t.Fatalf("Net Score is out of range")
// 		}
// 	}

	

// 	defer resp.Body.Close()
// }

<<<<<<<< HEAD:ratom/ratom_test.go
=======


		if net < 0 || net > 1 {
			t.Fatalf("Net Score is out of range")
		}
	}

	

	defer resp.Body.Close()
}

>>>>>>> 7d64c38 (Updating scoring and testing)
// Tests 6 - 12
func TestBusFactor2(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
<<<<<<< HEAD
========
// // Tests 6 - 12
// func TestBusFactor2(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go

// 	resp, error := httpClient.Get(endpoint2)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
		var bus = metrics.BusFactor(jsonRes)
========
// 		var bus = BusFactor(jsonRes)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go

// 		if bus < 0 || bus > 1 {
// 			t.Fatalf("Bus is out of range")
// 		}
// 	}

// 	defer resp.Body.Close()
// }

// func TestCorrectness2(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	resp, error := httpClient.Get(endpoint2)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
		var cor = metrics.Correctness(jsonRes)
========
// 		var cor = Correctness(jsonRes)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go


// 		if cor < 0 || cor > 1 {
// 			t.Fatalf("Correctness is out of range")
// 		}
// 	}

	

// 	defer resp.Body.Close()
// }

// func TestRampUp2(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	resp, error := httpClient.Get(endpoint2)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
		var ramp = metrics.RampUp(jsonRes, 70)
========
// 		var ramp = RampUp(jsonRes, 70)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go


// 		if ramp < 0 || ramp > 1 {
// 			t.Fatalf("Ramp Up is out of range")
// 		}
// 	}

// 	defer resp.Body.Close()
// }

// func TestResponsiveMaintainer2(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	resp, error := httpClient.Get(endpoint2)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
		var resp = metrics.ResponsiveMaintainer(jsonRes)
========
// 		var resp = ResponsiveMaintainer(jsonRes)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go

// 		if resp < 0 || resp > 1 {
// 			t.Fatalf("Responsive Maintainer is out of range")
// 		}
// 	}

	

// 	defer resp.Body.Close()
// }


// func TestNetScore2(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	resp, error := httpClient.Get(endpoint2)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
=======

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

		var bus = metrics.BusFactor(jsonRes)

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

		var cor = metrics.Correctness(jsonRes)


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

		var ramp = metrics.RampUp(jsonRes, 70)


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

		var resp = metrics.ResponsiveMaintainer(jsonRes)

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

>>>>>>> 7d64c38 (Updating scoring and testing)
		var bus = metrics.BusFactor(jsonRes)
		var cor = metrics.Correctness(jsonRes)
		var ramp = metrics.RampUp(jsonRes, 130)
		var resp = metrics.ResponsiveMaintainer(jsonRes)
		var net = metrics.NetScore(cor, bus, ramp, resp, false)
<<<<<<< HEAD
========
// 		var bus = BusFactor(jsonRes)
// 		var cor = Correctness(jsonRes)
// 		var ramp = RampUp(jsonRes, 130)
// 		var resp = ResponsiveMaintainer(jsonRes)
// 		var net = NetScore(cor, bus, ramp, resp, false)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go


// 		if net < 0 || net > 1 {
// 			t.Fatalf("Net Score is out of range")
// 		}
// 	}

	

// 	defer resp.Body.Close()
// }

<<<<<<<< HEAD:ratom/ratom_test.go
=======


		if net < 0 || net > 1 {
			t.Fatalf("Net Score is out of range")
		}
	}

	

	defer resp.Body.Close()
}

>>>>>>> 7d64c38 (Updating scoring and testing)
// Tests 13 - 18
func TestBusFactor3(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
<<<<<<< HEAD
========
// // Tests 13 - 18
// func TestBusFactor3(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go

// 	resp, error := httpClient.Get(endpoint3)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
		var bus = metrics.BusFactor(jsonRes)
========
// 		var bus = BusFactor(jsonRes)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go

// 		if bus < 0 || bus > 1 {
// 			t.Fatalf("Bus is out of range")
// 		}
// 	}

// 	defer resp.Body.Close()
// }

// func TestCorrectness3(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	resp, error := httpClient.Get(endpoint3)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
		var cor = metrics.Correctness(jsonRes)
========
// 		var cor = Correctness(jsonRes)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go


// 		if cor < 0 || cor > 1 {
// 			t.Fatalf("Correctness is out of range")
// 		}
// 	}

	

// 	defer resp.Body.Close()
// }

// func TestRampUp3(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	resp, error := httpClient.Get(endpoint3)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
		var ramp = metrics.RampUp(jsonRes, 600)
========
// 		var ramp = RampUp(jsonRes, 600)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go


// 		if ramp < 0 || ramp > 1 {
// 			t.Fatalf("Ramp Up is out of range")
// 		}
// 	}

// 	defer resp.Body.Close()
// }

// func TestResponsiveMaintainer3(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	resp, error := httpClient.Get(endpoint3)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
		var resp = metrics.ResponsiveMaintainer(jsonRes)
========
// 		var resp = ResponsiveMaintainer(jsonRes)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go

// 		if resp < 0 || resp > 1 {
// 			t.Fatalf("Responsive Maintainer is out of range")
// 		}
// 	}

	

// 	defer resp.Body.Close()
// }


// func TestNetScore3(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	resp, error := httpClient.Get(endpoint3)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
=======

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

		var bus = metrics.BusFactor(jsonRes)

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

		var cor = metrics.Correctness(jsonRes)


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

		var ramp = metrics.RampUp(jsonRes, 600)


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

		var resp = metrics.ResponsiveMaintainer(jsonRes)

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

>>>>>>> 7d64c38 (Updating scoring and testing)
		var bus = metrics.BusFactor(jsonRes)
		var cor = metrics.Correctness(jsonRes)
		var ramp = metrics.RampUp(jsonRes, 74)
		var resp = metrics.ResponsiveMaintainer(jsonRes)
		var net = metrics.NetScore(cor, bus, ramp, resp, false)
<<<<<<< HEAD
========
// 		var bus = BusFactor(jsonRes)
// 		var cor = Correctness(jsonRes)
// 		var ramp = RampUp(jsonRes, 74)
// 		var resp = ResponsiveMaintainer(jsonRes)
// 		var net = NetScore(cor, bus, ramp, resp, false)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go


// 		if net < 0 || net > 1 {
// 			t.Fatalf("Net Score is out of range")
// 		}
// 	}

// 	defer resp.Body.Close()
// }

<<<<<<<< HEAD:ratom/ratom_test.go
// Test 19
========
// // Test 19
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go

// func TestBusFactor_private(t *testing.T) {
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	resp, error := httpClient.Get(endpoint4)

// 	if error != nil {
// 		panic(error)
// 	}

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)

// 		if err != nil {
// 			panic(error)
// 		}
// 		bodyString := string(bodyBytes)

// 		// Citation needed for this
// 		resBytes := []byte(bodyString)
// 		var jsonRes map[string]interface{}
// 		_ = json.Unmarshal(resBytes, &jsonRes)

<<<<<<<< HEAD:ratom/ratom_test.go
		var bus = metrics.BusFactor(jsonRes)
========
// 		var bus = BusFactor(jsonRes)
>>>>>>>> 7d64c38 (Updating scoring and testing):ratom/metrics/metrics_test.go

// 		if bus < 0 || bus > 1 {
// 			t.Fatalf("Bus is out of range")
// 		}
// 	}

// 	defer resp.Body.Close()
// }
=======


		if net < 0 || net > 1 {
			t.Fatalf("Net Score is out of range")
		}
	}

	defer resp.Body.Close()
}

// Test 19

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

		var bus = metrics.BusFactor(jsonRes)

		if bus < 0 || bus > 1 {
			t.Fatalf("Bus is out of range")
		}
	}

	defer resp.Body.Close()
}
>>>>>>> 7d64c38 (Updating scoring and testing)
