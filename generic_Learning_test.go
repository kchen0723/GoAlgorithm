package main

import (
	"fmt"
	"testing"
)

func TestIterationArray(t *testing.T) {
	var s = []string{"foo", "bar", "zoo"}
	index := SliceIndex(s, "zoo")
	fmt.Println("index of zoo: ", index)
	index2 := SliceIndex[[]string, string](s, "zoo")
	fmt.Println(index2)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
