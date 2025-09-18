package dp

import (
	"math"
)

// bags: bag weight and profit if it is selected, for example, [2, 6] means weight is 2 and profit is 6
// maxWeight: max weight you can take, every item can take only once
// return max profit you can get
func GetMaxProfitForWeightByBackTracking(bags [][]float64, maxWeight float64) float64 {
	if len(bags) == 0 {
		return 0
	}
	selected := make([]bool, len(bags))
	profit := float64(0)
	getMaxProfitForWeightByBackTrackingHelp(bags, maxWeight, 0, &profit, selected)
	return profit
}

func getMaxProfitForWeightByBackTrackingHelp(bags [][]float64, maxWeight float64, index int, profit *float64, selected []bool) {
	totalWeight, totalProfit := float64(0), float64(0)
	for i, isSelected := range selected {
		if isSelected {
			totalWeight += bags[i][0]
			totalProfit += bags[i][1]
		}
	}
	if totalWeight <= maxWeight && totalProfit > *profit {
		*profit = totalProfit
	}
	if totalWeight >= maxWeight {
		return
	}

	for i := index; i < len(bags); i++ {
		if selected[i] {
			continue
		}

		selected[i] = true
		getMaxProfitForWeightByBackTrackingHelp(bags, maxWeight, i+1, profit, selected)
		selected[i] = false
	}
}

func GetMaxProfitForWeightByDP(bags [][]float64, maxWeight float64) float64 {
	if len(bags) == 0 {
		return 0
	}

	dp := make([][]float64, len(bags)+1) //each item means max profit for the first N bags for M weight. Add zero columns and zero rows.
	for i := 0; i < len(bags)+1; i++ {   //iterate all bags, the first row is zero
		dp[i] = make([]float64, int(maxWeight)+1)
		for j := 0; j < int(maxWeight)+1; j++ { //iterate all weights
			if i == 0 || j == 0 { //for first column and row, all are zeros.
				dp[i][j] = 0
			} else if j >= int(bags[i-1][0]) { //if weight is greater than candidate, we may take one bag
				taken := bags[i-1][1] + dp[i-1][j-int(bags[i-1][0])]
				notTaken := dp[i-1][j]
				dp[i][j] = math.Max(taken, notTaken)
			} else { //if weight is less than candidate, we cannot take any bag.
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(bags)][int(maxWeight)]
}
