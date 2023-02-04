package ratom

type Module struct {
	url         string
	netScore    float32
	rampUp      float32
	correctness int8
	busFactor   int8
	respMaint   float32
	license     bool
}

func Analyze(url string) Module {
	m := Module{url, -1, -1, -1, -1, -1, false}
	return m
}
