package main

import "strconv"

func diffWaysToComputer(input string) []float64 {
	var result []float64
	for i := 0; i < len(input); i++ {
		item := input[i]
		if item == '+' || item == '-' || item == '*' || item == '/' {
			leftString := input[:i]
			rightString := input[i+1:]
			leftResult := diffWaysToComputer(leftString)
			rightResult := diffWaysToComputer(rightString)

			for _, left := range leftResult {
				for _, right := range rightResult {
					switch item {
					case '+':
						result = append(result, left+right)
					case '-':
						result = append(result, left-right)
					case '*':
						result = append(result, left*right)
					case '/':
						if right != 0 {
							result = append(result, left/right)
						}
					}
				}
			}
		}
	}
	if len(result) == 0 {
		inputNumber, _ := strconv.ParseFloat(input, 64)
		result = append(result, inputNumber)
	}
	return result
}
