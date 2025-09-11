package tree

import (
	"testing"
)

func TestNewMaxHeapTreeByArray(t *testing.T) {
	var array = []float64{8, 7, 3, 6, 5, 9, 2, 4}
	tree := NewMaxHeapTreeByArray(array)
	println(tree.Length())
}
