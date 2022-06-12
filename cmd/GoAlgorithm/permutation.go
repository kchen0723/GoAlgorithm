package main

func GetAllPermutation(numbers []int, length int) [][]int {
	result := [][]int{}
	permutation := []int{}
	result = GetAllPermutationHelper(numbers, length, permutation, result)
	return result
}

func GetAllPermutationHelper(numbers []int, length int, permutation []int, result [][]int) [][]int {
	if len(permutation) == length {
		copyPermu := make([]int, len(permutation))
		copy(copyPermu, permutation)
		result = append(result, copyPermu)
		return result
	}

	for _, item := range numbers {
		permutation = append(permutation, item)
		result = GetAllPermutationHelper(numbers, length, permutation, result)
		permutation = permutation[:len(permutation)-1]
	}
	return result
}

func GetPermutation(numbers []int, length int) [][]int {
	result := [][]int{}
	permutation := []int{}
	used := make([]bool, len(numbers))
	result = GetPermutationHelper(numbers, length, used, permutation, result)
	return result
}

func GetPermutationHelper(numbers []int, length int, used []bool, permutation []int, result [][]int) [][]int {
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
		result = GetPermutationHelper(numbers, length, used, permutation, result)
		permutation = permutation[:len(permutation)-1]
		used[i] = false
	}
	return result
}

func GetSubsets(numbers []int, length int) [][]int {
	result := [][]int{}
	permutation := []int{}
	result = GetSubsetsHelper(numbers, 0, permutation, result)
	return result
}

func GetSubsetsHelper(numbers []int, start int, permutation []int, result [][]int) [][]int {
	// if len(permutation) == length {
	copyPermu := make([]int, len(permutation))
	copy(copyPermu, permutation)
	result = append(result, copyPermu)
	// return result
	// }

	for i := start; i < len(numbers); i++ {
		permutation = append(permutation, numbers[i])
		result = GetSubsetsHelper(numbers, i+1, permutation, result)
		permutation = permutation[:len(permutation)-1]
	}
	return result
}

func GetLengthSubset(numbers []int, length int) [][]int {
	result := [][]int{}
	permutation := []int{}
	result = GetLengthSubsetHelper(numbers, 0, length, permutation, result)
	return result
}

func GetLengthSubsetHelper(numbers []int, start int, length int, permutation []int, result [][]int) [][]int {
	if len(permutation) == length {
		copyPermu := make([]int, len(permutation))
		copy(copyPermu, permutation)
		result = append(result, copyPermu)
		return result
	}

	for i := start; i < len(numbers); i++ {
		permutation = append(permutation, numbers[i])
		result = GetLengthSubsetHelper(numbers, i+1, length, permutation, result)
		permutation = permutation[:len(permutation)-1]
	}
	return result
}
