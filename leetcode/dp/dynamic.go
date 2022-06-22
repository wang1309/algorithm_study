package dp

import "math"

// https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn854d/
// 爬楼梯
func climbStairs(n int) int {
	p, q, r := 0, 0, 1

	for i := 1; i < n; i++ {
		p = q
		q = r
		r = p + q
	}

	return r
}

// https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn8fsh/
// 买卖股票的最佳时机
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	minPrice := math.MaxInt64
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		if prices[i] > minPrice && prices[i] - minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}
