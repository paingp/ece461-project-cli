package metrics

func RampUp(jsonRes map[string]interface{}) float32 {
	wiki := 0.0
	pages := 0.0
	discussions := 0.0

	if jsonRes["has_wiki"].(bool) {
		wiki = .15
	}

	if jsonRes["has_pages"].(bool) {
		pages = .2
	}

	if jsonRes["has_discussions"].(bool) {
		discussions = .25
	}

	return float32(wiki + pages + discussions)
}