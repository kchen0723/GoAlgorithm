package main

import "testing"

func TestIterationArray(t *testing.T) {
	arr := []int{5, 8, 9}
	IterationArray(arr)
	RecursiveArray(arr)
}
