package metrics

import (
	//"fmt"

	"math"
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

	////////////////////////////////////////////////////////////////////////

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

	// Arbitrarily taken from the 15 of the month

	t1 := time.Date(year, monthObj, 15, 0, 0, 0, 0, time.UTC)
	t2 := time.Now()
	diff := t2.Sub(t1)

	var updatedLast float32
	//updatedLast = float32(diff.Seconds())

	//updatedLast = 0.25 * float32(math.Min(0, 1-math.Log10(float64(updatedLast))/8))
	if 0 < diff.Seconds() && diff.Seconds() <= 604800 { // 7 days timeline
		updatedLast = .25
	} else if diff.Seconds() <= 15720000 { // 1/2 a year timeline
		updatedLast = 0.12
	} else if diff.Seconds() <= 15720000*2 { // 1 year timeline
		updatedLast = 0.06
	} else if diff.Seconds() <= 15720000*2*2 { //2 years timeline
		updatedLast = 0.03
	} else {
		updatedLast = 0
	}

	////////////////////////////////////////////////////////////////////////

	hasIssues := jsonRes["has_issues"].(bool)
	//fmt.Println(hasIssues)

	openIssues := jsonRes["open_issues"].(float64)

	issuesScore := 0.0

	if hasIssues {
		issuesScore = 0.35 * math.Min(1, openIssues/350)
	}
	// if (hasIssues && (0 < openIssues && openIssues <= 50)){
	// 	issuesScore = .15
	// } else if (hasIssues && openIssues <= 100){
	// 	issuesScore = 0.2
	// } else if (hasIssues && openIssues <= 200){
	// 	issuesScore = 0.25
	// } else if (hasIssues && openIssues <= 300){
	// 	issuesScore = 0.3
	// } else if (hasIssues){
	// 	issuesScore = 0.35
	//}

	archivedStatus := jsonRes["archived"].(bool)
	archivedScore := 0.0

	if !archivedStatus {
		archivedScore = 0.2
	}

	totalValue := float32(private + updatedLast + float32(issuesScore) + float32(archivedScore))
	return totalValue
}
