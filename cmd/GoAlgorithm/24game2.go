package main

// use divide and conquer to calculate 24 game.

func calculate24() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	maxCards := 4
	Combinations := GetCombinationMultipleTimes(numbers, maxCards)
	operators := []int{0, 1, 2, 3}
	operatorPermutations := GetPermutationMultipleTimes(operators, maxCards-1)
	result := 0
	for _, item := range Combinations {
		result += calculateCombination(item, operatorPermutations)
	}
	if result > 0 {

	}
}

func calculateCombination(numbers []int, operators [][]int) int {
	result := 0
	permutations := GetPermutationUnique(numbers, len(numbers))
	for _, permutationNumbers := range permutations {
		for _, operators := range operators {
			candidates := tackleBrackets(permutationNumbers, operators)
			if hasTarget24(candidates) {
				return 1
			}
		}
	}
	return result
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
