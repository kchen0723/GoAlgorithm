package dp

import (
	"testing"
)

func TestGetMaxProfitForWeightByBackTracking(t *testing.T) {
	bags := [][]float64{{2, 6}, {2, 3}, {6, 5}, {5, 4}, {4, 6}}
	result := GetMaxProfitForWeightByBackTracking(bags, 10)
	if result != 15 {
		t.Errorf("GetMaxProfitForWeightByBackTracking() = %f, want %d", result, 15)
	}

	result = GetMaxProfitForWeightByDP(bags, 10)
	if result != 15 {
		t.Errorf("GetMaxProfitForWeightByBackTracking() = %f, want %d", result, 15)
	}
}
