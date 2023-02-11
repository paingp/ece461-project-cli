package metrics

import (
	"math"
)

func Correctness(jsonRes map[string]interface{}) float32 {

	ownerType := 0.0
	webCommit := 0.0

	stargazers := jsonRes["stargazers_count"].(float64)
	//watchers := jsonRes["watchers_count"].(float64)
	forksNum := jsonRes["forks_count"].(float64)
	//networkCount := jsonRes["network_count"].(float64)
	subscriberCount := jsonRes["subscribers_count"].(float64)

	owner_map := jsonRes["owner"].(map[string]interface{})
	if owner_map["type"].(string) == "Organization" {
		ownerType = .15
	} else {
		ownerType = .07
	}

	if jsonRes["web_commit_signoff_required"].(bool) {
		webCommit = .1
	} else {
		webCommit = 0.05
	}

	// Assigning weights to stargazers
	stargazers = math.Log10(stargazers) // Finding log of stargazers
	stargazers = stargazers / 105.5 // Normalizaing to 100,000
	stargazers = math.Min(1, stargazers) // Capping stargazers at 1
	// if 0 < stargazers && stargazers <= 50 {
	// 	stargazers = .05
	// } else if stargazers <= 200 {
	// 	stargazers = .1
	// } else if stargazers <= 1000 {
	// 	stargazers = .2
	// } else if stargazers <= 10000 {
	// 	stargazers = .25
	// } else {
	// 	stargazers = .3
	// }

	forksNum = math.Log10(forksNum) 
	forksNum = forksNum / 105.5
	forksNum = math.Min(1, forksNum)
	// if forksNum <= 50 {
	// 	forksNum = .05
	// } else if forksNum <= 200 {
	// 	forksNum = .1
	// } else if forksNum <= 1000 {
	// 	forksNum = .2
	// } else if forksNum <= 10000 {
	// 	forksNum = .25
	// } else {
	// 	forksNum = .3
	// }

	subscriberCount = math.Log10(subscriberCount)
	subscriberCount = subscriberCount / 105.5
	subscriberCount = math.Min(1, subscriberCount)
	// if subscriberCount <= 50 {
	// 	subscriberCount = .02
	// } else if subscriberCount <= 200 {
	// 	subscriberCount = .05
	// } else if subscriberCount <= 1000 {
	// 	subscriberCount = .1
	// } else if subscriberCount <= 10000 {
	// 	subscriberCount = .125
	// } else {
	// 	subscriberCount = .15
	// }

	total := math.Min(1, ownerType + webCommit + stargazers + forksNum + subscriberCount)
	return float32(total)
}
