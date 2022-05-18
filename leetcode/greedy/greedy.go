package greedy

// https://leetcode.cn/problems/candy/
// 分发糖果
func candy(ratings []int) (ans int) {
	n := len(ratings)
	left := make([]int, n)
	for i, r := range ratings {
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[i], right)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


// leetcode https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
// 买卖股票的最佳时机 II
// 7,1,5,3,6,4
// [1,2,3,4,5]
// [6,1,3,2,4,7]
func MaxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	purchasePrice := prices[0]
	profit := 0
	allProfit := 0

	for i:=1; i<len(prices); i++ {
		if 	purchasePrice < prices[i] && prices[i]-purchasePrice > profit {
			profit = prices[i]-purchasePrice
		} else {
			purchasePrice = prices[i]
			allProfit += profit
			profit = 0
		}
	}

	allProfit += profit

	return allProfit
}

// https://leetcode.cn/problems/remove-k-digits/
// 移掉 K 位数字
