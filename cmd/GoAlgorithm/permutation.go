package main

import "sort"

func GetPermutationUnique(numbers []int, length int) [][]int {
	result := [][]int{}
	permutation := []int{}
	used := make([]bool, len(numbers))
	result = GetPermutationUniqueHelper(numbers, length, used, permutation, result)
	return result
}

func GetPermutationUniqueHelper(numbers []int, length int, used []bool, permutation []int, result [][]int) [][]int {
	if len(permutation) == length {
		copyPermu := make([]int, len(permutation))
		copy(copyPermu, permutation)
		result = append(result, copyPermu)
		return result
	}

	for i := 0; i < len(numbers); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		permutation = append(permutation, numbers[i])
		result = GetPermutationUniqueHelper(numbers, length, used, permutation, result)
		permutation = permutation[:len(permutation)-1]
		used[i] = false
	}
	return result
}

func GetPermutationDuplicate(numbers []int, length int) [][]int {
	result := [][]int{}
	permutation := []int{}
	used := make([]bool, len(numbers))
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	result = GetPermutationDuplicateHelper(numbers, length, used, permutation, result)
	return result
}

func GetPermutationDuplicateHelper(numbers []int, length int, used []bool, permutation []int, result [][]int) [][]int {
	if len(permutation) == length {
		copyPermu := make([]int, len(permutation))
		copy(copyPermu, permutation)
		result = append(result, copyPermu)
		return result
	}

	for i := 0; i < len(numbers); i++ {
		if used[i] {
			continue
		}
		if i > 0 && numbers[i] == numbers[i-1] && used[i-1] {
			continue
		}

		used[i] = true
		permutation = append(permutation, numbers[i])
		result = GetPermutationDuplicateHelper(numbers, length, used, permutation, result)
		permutation = permutation[:len(permutation)-1]
		used[i] = false
	}
	return result
}

func GetPermutationMultipleTimes(numbers []int, length int) [][]int {
	result := [][]int{}
	permutation := []int{}
	result = GetPermutationMultipleTimesHelper(numbers, length, permutation, result)
	return result
}

func GetPermutationMultipleTimesHelper(numbers []int, length int, permutation []int, result [][]int) [][]int {
	if len(permutation) == length {
		copyPermu := make([]int, len(permutation))
		copy(copyPermu, permutation)
		result = append(result, copyPermu)
		return result
	}

	for i := 0; i < len(numbers); i++ {
		permutation = append(permutation, numbers[i])
		result = GetPermutationMultipleTimesHelper(numbers, length, permutation, result)
		permutation = permutation[:len(permutation)-1]
	}
	return result
}

func GetPermutationMultipleTimesSubset(numbers []string, length int) [][]string {
	var result [][]string
	var permutation []string
	result = GetPermutationMultipleTimesSubsetHelper(numbers, length, permutation, result)
	return result
}

func GetPermutationMultipleTimesSubsetHelper(numbers []string, length int, permutation []string, result [][]string) [][]string {
	copyPermu := make([]string, len(permutation))
	copy(copyPermu, permutation)
	result = append(result, copyPermu)
	if len(permutation) >= length {
		return result
	}

	for i := 0; i < len(numbers); i++ {
		permutation = append(permutation, numbers[i])
		result = GetPermutationMultipleTimesSubsetHelper(numbers, length, permutation, result)
		permutation = permutation[:len(permutation)-1]
	}
	return result
}
