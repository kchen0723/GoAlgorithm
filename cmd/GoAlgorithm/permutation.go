package main

func GetPermutation(numbers []int, length int) [][]int {
	result := [][]int{}
	permutation := []int{}
	result = GetPermutationHelper(numbers, length, permutation, result)
	return result
}

func GetPermutationHelper(numbers []int, length int, permutation []int, result [][]int) [][]int {
	if len(permutation) == length {
		copyPermu := make([]int, len(permutation))
		copy(copyPermu, permutation)
		result = append(result, copyPermu)
		return result
	}

	for _, item := range numbers {
		permutation = append(permutation, item)
		result = GetPermutationHelper(numbers, length, permutation, result)
		permutation = permutation[:len(permutation)-1]
	}
	return result
}
