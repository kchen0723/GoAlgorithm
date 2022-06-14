package main

import "sort"

func GetCombinationUnique(numbers []int, length int) [][]int {
	result := [][]int{}
	combination := []int{}
	result = GetCombinationUniqueHelper(numbers, 0, length, combination, result)
	return result
}

func GetCombinationUniqueHelper(numbers []int, start int, length int, combination []int, result [][]int) [][]int {
	if len(combination) == length {
		copyPermu := make([]int, len(combination))
		copy(copyPermu, combination)
		result = append(result, copyPermu)
		return result
	}

	for i := start; i < len(numbers); i++ {
		combination = append(combination, numbers[i])
		result = GetCombinationUniqueHelper(numbers, i+1, length, combination, result)
		combination = combination[:len(combination)-1]
	}
	return result
}

func GetCombinationDuplicate(numbers []int, length int) [][]int {
	result := [][]int{}
	combination := []int{}
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	result = GetCombinationDuplicateHelper(numbers, 0, length, combination, result)
	return result
}

func GetCombinationDuplicateHelper(numbers []int, start int, length int, combination []int, result [][]int) [][]int {
	if len(combination) == length {
		copyPermu := make([]int, len(combination))
		copy(copyPermu, combination)
		result = append(result, copyPermu)
		return result
	}

	for i := start; i < len(numbers); i++ {
		if i > start && numbers[i] == numbers[i-1] {
			continue
		}

		combination = append(combination, numbers[i])
		result = GetCombinationDuplicateHelper(numbers, i+1, length, combination, result)
		combination = combination[:len(combination)-1]
	}
	return result
}

func GetCombinationMultipleTimes(numbers []int, length int) [][]int {
	result := [][]int{}
	combination := []int{}
	result = GetCombinationMultipleTimesHelper(numbers, 0, length, combination, result)
	return result
}

func GetCombinationMultipleTimesHelper(numbers []int, start int, length int, combination []int, result [][]int) [][]int {
	if len(combination) == length {
		copyPermu := make([]int, len(combination))
		copy(copyPermu, combination)
		result = append(result, copyPermu)
		return result
	}

	for i := start; i < len(numbers); i++ {
		combination = append(combination, numbers[i])
		result = GetCombinationMultipleTimesHelper(numbers, i, length, combination, result)
		combination = combination[:len(combination)-1]
	}
	return result
}
