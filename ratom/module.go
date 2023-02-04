package ratom

type Module struct {
	url string
	netScore float32
	rampUp float32
	correctness uint8
	busFactor uint8	
	respMaint float32
	license bool
}

// func (url string) Analyze() Module {
// 	m := Module{url}
// 	return m
// }
