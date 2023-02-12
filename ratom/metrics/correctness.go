package metrics

import (
	"math"
)

// This function uses rest and graphQL through the GITHUB API 
// to collect data pertaining to the Correctness factor, then analyzes
// the data and returns a weighted sum of the scores
func Correctness(jsonRes map[string]interface{}) float32 {

	ownerType := 0.0
	webCommit := 0.0

	// Collecting data from API
	stargazers := jsonRes["stargazers_count"].(float64)
	forksNum := jsonRes["forks_count"].(float64)
	subscriberCount := jsonRes["subscribers_count"].(float64)

	// Analysis of owner type
	owner_map := jsonRes["owner"].(map[string]interface{})
	if owner_map["type"].(string) == "Organization" {
		ownerType = .15
	} else {
		ownerType = .07
	}

	// Analysis of web_commit_signoff_required
	if jsonRes["web_commit_signoff_required"].(bool) {
		webCommit = .1
	} else {
		webCommit = 0.05
	}

	// Assigning weights to stargazers
	stargazers = math.Log10(stargazers) // Finding log of stargazers
	stargazers = stargazers / 25
	stargazers = math.Min(1, stargazers) // Capping stargazers at 1

	// Assigning weights to forks
	forksNum = math.Log10(forksNum) 
	forksNum = forksNum / 25
	forksNum = math.Min(1, forksNum)

	// Assigning weights to subscriber count
	subscriberCount = math.Log10(subscriberCount)
	subscriberCount = subscriberCount / 25
	subscriberCount = math.Min(1, subscriberCount)

	total := math.Min(1, ownerType + webCommit + stargazers + forksNum + subscriberCount)
	return float32(total)
}
