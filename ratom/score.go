package ratom

import (
	"github.com/paingp/ece461-project-cli/ratom/metrics"
	"net/http"
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
	m := Module{url, -1, -1, c, -1, -1, false}
	return m
}
