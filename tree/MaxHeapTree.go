package tree

import (
	"math"
)

type MaxHeapTree struct {
	maxHeap []float64
}

func NewMaxHeapTree() *MaxHeapTree {
	return &MaxHeapTree{}
}

func NewMaxHeapTreeByArray(array []float64) *MaxHeapTree {
	result := &MaxHeapTree{
		maxHeap: array,
	}
	middle := len(array)/2 - 1
	for i := middle; i >= 0; i-- {
		result.heapify(i)
	}
	return result
}

func (mht *MaxHeapTree) Length() int {
	return len(mht.maxHeap)
}

func (mht *MaxHeapTree) Insert(value float64) {
	mht.maxHeap = append(mht.maxHeap, value)
	index := mht.Length() - 1
	half := (index - 1) / 2
	for mht.maxHeap[index] > mht.maxHeap[half] {
		temp := mht.maxHeap[index]
		mht.maxHeap[index] = mht.maxHeap[half]
		mht.maxHeap[half] = temp

		index = half
		half = (index - 1) / 2
	}
}

func (mht *MaxHeapTree) Pop() float64 {
	if mht.Length() == 0 {
		return math.NaN()
	}
	result := mht.maxHeap[0]
	mht.maxHeap[0] = mht.maxHeap[mht.Length()-1]
	mht.maxHeap = mht.maxHeap[:mht.Length()-1]
	mht.heapify(0)
	return result
}

func (mht *MaxHeapTree) heapify(index int) {
	if index < 0 || index >= mht.Length() {
		return
	}
	for {
		leftIndex := 2*index + 1
		rightIndex := 2*index + 2
		largerIndex := index
		if leftIndex < mht.Length() && mht.maxHeap[leftIndex] > mht.maxHeap[largerIndex] {
			largerIndex = leftIndex
		}
		if rightIndex < mht.Length() && mht.maxHeap[rightIndex] > mht.maxHeap[largerIndex] {
			largerIndex = rightIndex
		}

		if largerIndex == index {
			break
		}

		temp := mht.maxHeap[index]
		mht.maxHeap[index] = mht.maxHeap[largerIndex]
		mht.maxHeap[largerIndex] = temp

		index = largerIndex
	}
}
