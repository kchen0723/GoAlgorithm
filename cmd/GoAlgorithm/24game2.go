package main

// use divide and conquer to calculate 24 game.

type game struct {
	permutations []int
	operators    []int
	brackets     []int
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
			candidates := tackleBrackets(permutationNumbers, operators)
			if hasTarget24(candidates) {
				result.permutations = permutationNumbers
				result.operators = operators
				return true, result
			}
		}
	}
	return false, result
}

func tackleBrackets(numbers []int, operators []int) []float64 {
	var result []float64
	for i := 0; i < len(operators); i++ {
		leftNumbers := numbers[:i+1]
		rightNumbers := numbers[i+1:]
		leftOperators := operators[:i]
		rightOperators := operators[i+1:]

		leftResult := tackleBrackets(leftNumbers, leftOperators)
		rightResult := tackleBrackets(rightNumbers, rightOperators)
		for _, left := range leftResult {
			for _, right := range rightResult {
				switch operators[i] {
				case 0:
					result = append(result, left+right)
				case 1:
					result = append(result, left-right)
				case 2:
					result = append(result, left*right)
				case 3:
					if right != 0 {
						result = append(result, left/right)
					}
				}
			}
		}
	}

	if len(result) == 0 {
		result = append(result, float64(numbers[0]))
	}
	return result
}
