package main

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

func get4NumberResultHelper(numbers []int) bool {
	var result bool
	if len(numbers) != 4 {
		return result
	}

	return result
}
