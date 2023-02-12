package metrics

import (
	"fmt"
	"runtime"
	"strings"
)

var Functions []string

// Auxillary function used for logging (Cited in Report)
func File_line() string {
	_, fileName, fileLine, ok := runtime.Caller(1)
	var s string
	if ok {
		s = fmt.Sprintf("%s:%d", fileName, fileLine)
	} else {
		s = ""
	}

	fileInfo := strings.Split(s, ":")
	fileInfoSize := len(fileInfo)
	return fileInfo[fileInfoSize-1]
}

// Combining the weighted sum of the analyzed factors
func NetScore(correctness float32, busFactor float32, rampUp float32, responsiveness float32, license bool) float32 {
	return float32(.35*correctness + .25*responsiveness + .2*busFactor + .2*rampUp)
}
