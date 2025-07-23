package simple_algos

func MinimizeCoinChange(coins []int, target int) int {
	if target == 0 {
		return 0
	}
	var dp = make([]int, target+1)
	for i := range dp {
		dp[i] = target + 1
	}
	dp[0] = 0
	for amount := 1; amount <= target; amount++ {
		for _, coin := range coins {
			if coin <= amount {
				dp[amount] = min(dp[amount], 1+dp[amount-coin])
			}
		}
	}
	return dp[target]
}
