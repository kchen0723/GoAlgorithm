package slidewindow

func GetMinCover(source string, target string) string {
	if len(source) == 0 || len(target) == 0 {
		return ""
	}
	left, right, start, matchCount, length := 0, 0, 0, 0, len(source)
	matchWindow := make(map[rune]int)
	targetDictionary := buildDictionary(target)
	sourceRunes := []rune(source)
	for ; right < len(sourceRunes); right++ {
		item := sourceRunes[right]
		if _, ok := targetDictionary[item]; ok {
			matchWindow[item]++
			if matchWindow[item] == targetDictionary[item] {
				matchCount++
			}
		}

		for matchCount == len(targetDictionary) {
			if right-left+1 < length {
				start = left
				length = right - left + 1
			}

			removeCandidate := sourceRunes[left]
			left++
			if _, ok := targetDictionary[removeCandidate]; ok {
				if matchWindow[removeCandidate] == targetDictionary[removeCandidate] {
					matchCount--
				}
				matchWindow[removeCandidate]--
			}
		}
	}
	if length < len(source) {
		return string(sourceRunes[start : start+length])
	}
	return ""
}

func buildDictionary(target string) map[rune]int {
	result := make(map[rune]int)
	for _, item := range target {
		result[item]++
	}
	return result
}
