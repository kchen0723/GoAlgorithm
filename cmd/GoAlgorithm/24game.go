package main

import (
	"math"
)

//https://blog.csdn.net/problc/article/details/7287138, total 1362 answers
var target24 = 24

func get2NumberResult(left float64, right float64) []float64 {
	var result []float64
	result = append(result, left+right)
	result = append(result, left-right)
	result = append(result, left*right)
	if right != 0 {
		result = append(result, left/right)
	}
	return result
}

func Is4NumberCombationValid(numbers []int) bool {
	var result bool
	if len(numbers) != 4 {
		return result
	}

	// testNumbers := []int{1, 2, 3, 4}
	// permutations := GetPermutationUnique(testNumbers, len(testNumbers))
	permutations := GetPermutationUnique(numbers, len(numbers))
	for _, permutation := range permutations {
		if Is4NumberPermutationValid(permutation) {
			result = true
			break
		}
	}
	return result
}

func getOperatorPermutations(length int) [][]string {
	operators := []string{"+", "-", "*", "/"}
	operatorsPermutation := GetPermutationMultipleTimesSubset(operators, length)
	return operatorsPermutation
}

func get4NumberPostfix(numbers []int) [][]interface{} {
	var result [][]interface{}
	if len(numbers) != 4 {
		return result
	}

	var permuation []interface{}
	permuation = append(permuation, numbers[0])
	permuation = append(permuation, numbers[1])
	// oneOperatorPermutation := getOperatorPermutations(1)
	// twoeOperatorPermutation := getOperatorPermutations(2)
	// threeOperatorPermutation := getOperatorPermutations(3)
	// for _, one := range oneOperatorPermutation {

	// 	for _, two := range twoeOperatorPermutation {
	// 		for _, tree := range threeOperatorPermutation {

	// 		}
	// 	}
	// }

	return result
}

func Is4NumberPermutationValid(numbers []int) bool {
	if isCaseAValid(numbers) { //（（0， 1）2，）3  ab*c*d*
		return true
	}
	if isCaseBvalid(numbers) { //（0， 1）（2，3）ab*cd**
		return true
	}
	if isCaseCValid(numbers) { // 0（1（2，3））abcd***
		return true
	}
	if isCaseDValid(numbers) { //（0， (1, 2）)3  abc**d*
		return true
	}
	if isCaseEValid(numbers) { //0， ((1 2)，3) abc*d**
		return true
		//ab*                                        ab*
		//ab*c* abc**                                ab*c**
		//ab*c*d* ab*cd** abc*d** abcd*** abc**d*    ab*c**d***
		//ab*c*d*e* abc**d*e* abcde**** abc*d**e*    ab*c**d***e****
		//                                           ab*c**d***e****f*****
	}
	return false
}

func isCaseAValid(numbers []int) bool {
	if len(numbers) != 4 {
		return false
	}
	left2 := get2NumberResult(float64(numbers[0]), float64(numbers[1]))
	for _, item := range left2 {
		left3 := get2NumberResult(item, float64(numbers[2]))
		for _, left3item := range left3 {
			result := get2NumberResult(left3item, float64(numbers[3]))
			if hasTarget24(result) {
				return true
			}
		}
	}
	return false
}

func hasTarget24(numbers []float64) bool {
	for _, item := range numbers {
		// if item == float64(target24) {
		// 	return true
		// }
		if math.Abs(item-float64(target24)) < 0.001 {
			return true
		}
	}
	return false
}

func isCaseBvalid(numbers []int) bool {
	if len(numbers) != 4 {
		return false
	}
	left2 := get2NumberResult(float64(numbers[0]), float64(numbers[1]))
	right2 := get2NumberResult(float64(numbers[2]), float64(numbers[3]))
	for _, left := range left2 {
		for _, right := range right2 {
			result := get2NumberResult(left, right)
			if hasTarget24(result) {
				return true
			}
		}
	}
	return false
}

func isCaseCValid(numbers []int) bool {
	if len(numbers) != 4 {
		return false
	}
	right2 := get2NumberResult(float64(numbers[2]), float64(numbers[3]))
	for _, item := range right2 {
		right3 := get2NumberResult(float64(numbers[1]), item)
		for _, left3item := range right3 {
			result := get2NumberResult(float64(numbers[0]), left3item)
			if hasTarget24(result) {
				return true
			}
		}
	}
	return false
}

func isCaseDValid(numbers []int) bool {
	if len(numbers) != 4 {
		return false
	}
	middle2 := get2NumberResult(float64(numbers[1]), float64(numbers[2]))
	for _, item := range middle2 {
		left3 := get2NumberResult(float64(numbers[0]), item)
		for _, left3item := range left3 {
			result := get2NumberResult(float64(numbers[3]), left3item)
			if hasTarget24(result) {
				return true
			}
		}
	}
	return false
}

func isCaseEValid(numbers []int) bool {
	if len(numbers) != 4 {
		return false
	}
	middle2 := get2NumberResult(float64(numbers[1]), float64(numbers[2]))
	for _, item := range middle2 {
		left3 := get2NumberResult(item, float64(numbers[3]))
		for _, left3item := range left3 {
			result := get2NumberResult(float64(numbers[0]), left3item)
			if hasTarget24(result) {
				return true
			}
		}
	}
	return false
}
