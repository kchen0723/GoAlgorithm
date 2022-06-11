package main

var result [][]int

func GetPermutation(numbers []int, length int) [][]int {
	result = [][]int{}
	permutation := []int{}
	GetPermutationHelper(numbers, length, permutation)
	return result
}

func GetPermutationHelper(numbers []int, length int, permutation []int) {
	if len(permutation) == length {
		copyPermu := make([]int, len(permutation))
		copy(copyPermu, permutation)
		result = append(result, copyPermu)
		return
	}

	for _, item := range numbers {
		permutation = append(permutation, item)
		GetPermutationHelper(numbers, length, permutation)
		permutation = permutation[:len(permutation)-1]
	}
}
