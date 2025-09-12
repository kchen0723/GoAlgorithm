package dp

//bags: bag weight and profit if it is selected, for example, [2, 6] means weight is 2 and profit is 6
//maxWeight: max weight you can take, every item can take only once
//return max profit you can get
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
