package metrics

func NetScore(correctness float32, busFactor float32, rampUp float32, responsiveness float32, license bool) float32 {

	return float32(.35*correctness + .25*responsiveness + .2*busFactor + .2*rampUp)
}
