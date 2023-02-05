package metrics

func BusFactor(jsonRes map[string]interface{}) float32 {

	var disabled float32
	var forking float32
	var visibility float32
	
	if jsonRes["web_commit_signoff_required"].(bool) {
		disabled = .0
	} else {
		disabled = 0.2
	}
	
	if jsonRes["web_commit_signoff_required"].(bool) {
		forking = .0
	} else {
		forking = 0.4
	}
	
	if jsonRes["visibility"].(string) == "public" {
		visibility = .4
	} else {
		visibility = .2
	}
	// disabled
	// allow_forking
	// visibility
	return float32(disabled + forking + visibility)
}
