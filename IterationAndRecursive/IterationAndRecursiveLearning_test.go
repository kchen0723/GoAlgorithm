package main

import "testing"

func TestGetMinimumPostOfficeDistance(t *testing.T) {
	arr := []int{0, 1, 2, 3}
	actual := GetMinimumPostOfficeDistance(arr, 5)
	if actual != 0 {
		t.Error("the post office distance is wrong")
	}
}
