package tree

type MaxHeapTree struct {
	MaxHeap []float64
}

func NewMaxHeapTree() *MaxHeapTree {
	return &MaxHeapTree{}
}

func NewMaxHeapTreeByArray(array []float64) *MaxHeapTree {
	result := &MaxHeapTree{
		MaxHeap: array,
	}
	middle := len(array)/2 - 1
	for i := middle; i >= 0; i-- {
		result.heapify(i)
	}
	return result
}

func (mht *MaxHeapTree) Length() int {
	return len(mht.MaxHeap)
}

func (mht *MaxHeapTree) heapify(index int) {
	if index < 0 || index >= mht.Length() {
		return
	}
	for {
		leftIndex := 2*index + 1
		rightIndex := 2*index + 2
		largerIndex := index
		if leftIndex < mht.Length() && mht.MaxHeap[leftIndex] > mht.MaxHeap[largerIndex] {
			largerIndex = leftIndex
		}
		if rightIndex < mht.Length() && mht.MaxHeap[rightIndex] > mht.MaxHeap[largerIndex] {
			largerIndex = rightIndex
		}

		if largerIndex == index {
			break
		}

		temp := mht.MaxHeap[index]
		mht.MaxHeap[index] = mht.MaxHeap[largerIndex]
		mht.MaxHeap[largerIndex] = temp

		index = largerIndex
	}
}
