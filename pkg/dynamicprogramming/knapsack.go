package dynamicprogramming

import "data-structures/pkg/utils"

/*
Given two integer arrays to represent weights and profits of ‘N’ items, we need to find a subset of these items which
will give us maximum profit such that their cumulative weight is not more than a given number ‘C’.
Write a function that returns the maximum profit.
Each item can only be selected once, which means either we put an item in the knapsack or skip it.
*/

// KnapsackWithMemoization calculate the maximum profit using recursion and memoization technique, it is a top-dow approach
func KnapsackWithMemoization(profits []int, weights []int, capacity int) int {

	var dp func(idx, cap int) int
	memo := utils.Make2dArray(len(weights), capacity + 1)

	dp = func (idx, cap int) int {
		if cap <= 0 || idx >= len(profits) {
			return 0
		}

		if memo[idx][cap] != 0 {
			return memo[idx][cap]
		}

		profitIncludeCurrentIdx := 0
		if weights[idx] <= cap {
			profitIncludeCurrentIdx = profits[idx] + dp(idx + 1, cap - weights[idx])
		}

		profitWithoutCurrentIdx := dp(idx + 1, cap)
		memo[idx][cap] = utils.Max(profitIncludeCurrentIdx, profitWithoutCurrentIdx)

		return memo[idx][cap]

	}

	return dp(0, capacity)
}

// Knapsack calculate the maximum profit using recursion and brute force, it is top-down approach
func Knapsack(profits []int, weights []int, capacity int) int {

	var dp func(idx, cap int) int

	dp = func (idx, cap int) int {
		if cap <= 0 || idx >= len(profits) {
			return 0
		}

		profitIncludeCurrentIdx := 0
		if weights[idx] <= cap {
			profitIncludeCurrentIdx = profits[idx] + dp(idx + 1, cap - weights[idx])
		}

		profitWithoutCurrentIdx := dp(idx + 1, cap)

		return utils.Max(profitIncludeCurrentIdx, profitWithoutCurrentIdx)

	}

	return dp(0, capacity)
}

// KnapsackWithTabulation calculate the maximum profit using tabulation, it is a bottom-up approach
func KnapsackWithTabulation(profits []int, weights []int, capacity int) int {
	dp := utils.Make2dArray(len(weights), capacity + 1)

	for i := 0; i < len(weights); i++ {
		for c := 0; c < capacity; c++ {
			dp[i][c] = utils.Max (dp[i-1][c], profits[i] + dp[i-1][c-weights[i]])
		}
	}

	return dp[len(weights)][capacity]
}
