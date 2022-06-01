package main

// import "strings"

// func GetTokens(source string, splitors []string) []string {
// 	var result []string
// 	trimmedLine := strings.TrimSpace(source)
// 	trimmedLineLength := len(trimmedLine)
// 	if trimmedLineLength == 0 {
// 		return result
// 	}

// 	splitorsMap := buildSplitorMap(splitors)
// 	startPosition := 0
// 	for startPosition < trimmedLineLength {
// 		oneCharacter := string(source[startPosition])
// 		_, isExisted := splitorsMap[oneCharacter]
// 		if isExisted {
// 			result = append(result, oneCharacter)
// 			startPosition++
// 		} else {
// 			for i := 0; i < len(source)-startPosition; i++ {
// 				candidateString := source[startPosition : startPosition+i]
// 				_, ok := splitorsMap[candidateString]
// 				if ok {
// 					result = append(result, candidateString)
// 					startPosition += i
// 				} else if i == len(source)-startPosition {

// 				}
// 			}
// 		}
// 	}
// 	return result
// }

// func buildSplitorMap(splitors []string) map[string]bool {
// 	result := make(map[string]bool)
// 	for _, item := range splitors {
// 		_, isExisted := result[item]
// 		if !isExisted {
// 			result[item] = true
// 		}
// 	}
// 	return result
// }
