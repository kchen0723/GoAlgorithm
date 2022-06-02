package main

import (
	"regexp"
)

//source := "(-5.3+72/9-6)*(7+3)"
func GetTokens(source string) []string {
	re := regexp.MustCompile(`[0-9]+[\.0-9]*|[\(\)\+\*-/]`)
	splits := re.FindAllStringSubmatch(source, -1)
	var result []string
	for _, item := range splits {
		if len(item) > 0 {
			result = append(result, item[0])
		}
	}
	return result
}
