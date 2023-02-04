package ratom

type Module struct {
	string url,
	int rampUp,
	int correctness,
	int busFactor,
	int respMaint,
	bool license
}

func (url string) analyze() Module {
	m := Module{url}
	return m
}
