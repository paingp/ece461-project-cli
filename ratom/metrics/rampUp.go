package metrics

func RampUp(jsonRes map[string]interface{}, totalComments int) float32 {
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

	var commentsScore float32

	if totalComments >= 0 && totalComments <= 10{
		commentsScore = 0.1
	} else if totalComments <= 50{
		commentsScore = 0.2
	} else if totalComments <= 100{
		commentsScore = 0.25
	} else if commentsScore <= 400 {
		commentsScore = 0.325
	} else {
		commentsScore = 0.4
	}

	return float32(wiki + pages + discussions + float64(commentsScore))
}
