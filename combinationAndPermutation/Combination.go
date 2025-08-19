package combinationandpermation

func GetCombationFromUniqueArray(nums []int, length int) [][]int {
	if len(nums) == 0 || length == 0 || length > len(nums) {
		return [][]int{}
	}
	return getCombationFromUniqueArrayHelp(nums, length, 0, []int{})
}

func getCombationFromUniqueArrayHelp(nums []int, length int, index int, candidate []int) [][]int {
	if len(candidate) == length {
		item := append([]int(nil), candidate...)
		return [][]int{item}
	}
	var result [][]int
	for i := index; i < len(nums); i++ {
		candidate = append(candidate, nums[i])
		items := getCombationFromUniqueArrayHelp(nums, length, i+1, candidate)
		result = append(result, items...)
		candidate = candidate[:len(candidate)-1]
	}
	return result
}
