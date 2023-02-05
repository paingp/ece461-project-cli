package metrics

import (
	//"fmt"
	"strconv"
	"strings"
	"time"
)

func ResponsiveMaintainer(jsonRes map[string]interface{}) float32 {

	var private float32

	updatedAt := jsonRes["updated_at"].(string)

	if jsonRes["private"].(bool) {
		private = .1
	} else {
		private = .05
	}

	updateDateList := strings.Split(updatedAt, "-")
	yearStr := updateDateList[0]
	monthStr := updateDateList[1]
	//dayStr := updateDateList[2]

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		panic(err)
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		panic(err)
	}
	monthObj := time.Month(month)

	t1 := time.Date(year, monthObj, 15, 0, 0, 0, 0, time.UTC)
	t2 := time.Now()
	diff := t2.Sub(t1)
	// diff.Seconds()

	var updatedLast float32

	if 0 < diff.Seconds() && diff.Seconds() <= 604800 { // 7 days timeline
		updatedLast = .25
	} else if diff.Seconds() <= 15720000 { // 1/2 a year timeline 
		updatedLast = 0.12
	} else if diff.Seconds() <= 15720000 * 2 { // 1 year timeline 
		updatedLast = 0.06
	} else if diff.Seconds() <= 15720000 * 2 * 2 { //2 years timeline 
		updatedLast = 0.03
	} else {
		updatedLast = 0
	}

	return private + updatedLast
}
