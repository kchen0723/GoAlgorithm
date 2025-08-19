package combinationandpermation

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func sortSlices(slices [][]int) {
	//sort inner slice
	for _, slice := range slices {
		sort.Ints(slice)
	}

	//sort the outer slice based on the inner slice
	sort.Slice(slices, func(i, j int) bool {
		for index := 0; index < len(slices[i]) && index < len(slices[j]); index++ {
			if slices[i][index] != slices[j][index] {
				return slices[i][index] < slices[j][index]
			}
		}
		return len(slices[i]) < len(slices[j])
	})
}

func TestGetCombationFromUniqueArray(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		length   int
		expected [][]int
	}{
		{
			name:     "stand case",
			nums:     []int{1, 2, 3},
			length:   2,
			expected: [][]int{{1, 2}, {1, 3}, {2, 3}},
		},
		{
			name:     "Length is 1",
			nums:     []int{1, 2, 3, 4},
			length:   1,
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name:     "Length equals array size",
			nums:     []int{1, 2, 3},
			length:   3,
			expected: [][]int{{1, 2, 3}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetCombationFromUniqueArray(tc.nums, tc.length)
			sortSlices(result)
			sortSlices(tc.expected)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("GetcombationFromUniqueArray(%v, %d) = %v, want %v", tc.nums, tc.length, result, tc.expected)
			}
		})
	}
}

func TestGetCombationFromDuplicateArray(t *testing.T) {
	nums := []int{1, 2, 3, 2}
	result := GetCombationFromDuplicateArray(nums, 2)
	fmt.Println(result)
}
