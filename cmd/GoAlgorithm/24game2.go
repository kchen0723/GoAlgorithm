package main

import (
	"math"
	"strconv"
)

// use divide and conquer to calculate 24 game.

type game struct {
	permutations []int
	operators    []int
	express      string
}

func calculate24() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	maxCards := 4
	Combinations := GetCombinationMultipleTimes(numbers, maxCards)
	operators := []int{0, 1, 2, 3}
	operatorPermutations := GetPermutationMultipleTimes(operators, maxCards-1)
	var games []game
	for _, item := range Combinations {
		ok, game := calculateCombination(item, operatorPermutations)
		if ok {
			games = append(games, game)
		}
	}
	if len(games) > 0 {

	}
}

func calculateCombination(numbers []int, operators [][]int) (bool, game) {
	var result game
	permutations := GetPermutationUnique(numbers, len(numbers))
	for _, permutationNumbers := range permutations {
		for _, operators := range operators {
			candidates, candidateExpresses := tackleBrackets(permutationNumbers, operators)
			has, index := getTarget24Index(candidates)
			if has {
				result.permutations = permutationNumbers
				result.operators = operators
				result.express = candidateExpresses[index]
				return true, result
			}
		}
	}
	return false, result
}

func getTarget24Index(numbers []float64) (bool, int) {
	for i, item := range numbers {
		// if item == float64(target24) {
		// 	return true
		// }
		if math.Abs(item-float64(target24)) < 0.001 {
			return true, i
		}
	}
	return false, 0
}

func tackleBrackets(numbers []int, operators []int) ([]float64, []string) {
	var result []float64
	var expressResult []string
	for i := 0; i < len(operators); i++ {
		leftNumbers := numbers[:i+1]
		rightNumbers := numbers[i+1:]
		leftOperators := operators[:i]
		rightOperators := operators[i+1:]

		leftResult, leftExpress := tackleBrackets(leftNumbers, leftOperators)
		rightResult, rightExpress := tackleBrackets(rightNumbers, rightOperators)
		for j, left := range leftResult {
			for k, right := range rightResult {
				switch operators[i] {
				case 0:
					result = append(result, left+right)
					expressResult = append(expressResult, "("+leftExpress[j]+" + "+rightExpress[k]+")")
				case 1:
					result = append(result, left-right)
					expressResult = append(expressResult, "("+leftExpress[j]+" - "+rightExpress[k]+")")
				case 2:
					result = append(result, left*right)
					expressResult = append(expressResult, "("+leftExpress[j]+" * "+rightExpress[k]+")")
				case 3:
					if right != 0 {
						result = append(result, left/right)
						expressResult = append(expressResult, "("+leftExpress[j]+" / "+rightExpress[k]+")")
					}
				}
			}
		}
	}

	if len(result) == 0 {
		result = append(result, float64(numbers[0]))
		expressResult = append(expressResult, strconv.Itoa(numbers[0]))
	}
	return result, expressResult
}
